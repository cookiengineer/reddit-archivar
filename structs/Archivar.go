package structs

import "reddit-archivar/schemas"
import "encoding/json"
import "strconv"
import "time"

type ArchivarTask struct {
	Subreddit string // "/r/Subreddit"
	Type      string // "hot" || "top" || "new" || "comments" || "search"
	Query     string // "search-keyword"
	Page      string // "t3_12345"
	Count     int
}

func (task *ArchivarTask) ToURL() string {

	var url string

	if task.Type == "new" || task.Type == "hot" || task.Type == "top" {

		url = "https://old.reddit.com/r/" + task.Subreddit + "/" + task.Type + ".json"
		url += "?t=all"
		url += "&limit=100"

		if task.Count != 0 {
			url += "&count=" + strconv.Itoa(task.Count)
		}

		if task.Page != "" {
			url += "&after=" + task.Page
		}

	} else if task.Type == "comments" {

		url = "https://old.reddit.com/r/" + task.Subreddit + "/comments/" + task.Page + ".json"

	} else if task.Type == "search" {

		url = "https://old.reddit.com/r/" + task.Subreddit + "/search.json"
		url += "?q=" + task.Query
		url += "&restrict_sr=on"
		url += "&t=all"

		if task.Count != 0 {
			url += "&count=" + strconv.Itoa(task.Count)
		}

		if task.Page != "" {
			url += "&after=" + task.Page
		}

	}

	return url

}

type Archivar struct {
	Cache     Cache
	Busy      bool
	Scraper   Scraper
	Subreddit string
	Errors    []string
	Tasks     []ArchivarTask
}

func processTasks(archivar *Archivar) {

	var filtered []ArchivarTask
	var limit int = archivar.Scraper.Limit

	if archivar.Scraper.Throttled == true {
		limit = 1
	}

	for t := 0; t < len(archivar.Tasks); t++ {

		if len(filtered) < limit {
			filtered = append(filtered, archivar.Tasks[t])
		} else {
			break
		}

	}

	if len(filtered) > 0 {

		for f := 0; f < len(filtered); f++ {

			var task = filtered[f]

			archivar.Scraper.DeferRequest(task.ToURL(), func(buffer []byte, status int) {

				var schema schemas.ThreadListing

				err := json.Unmarshal(buffer, &schema)

				if err == nil {
					processThreadListing(archivar, schema, task)
				}

			})

		}

		for archivar.Scraper.Busy == true {
			// Wait for Scraper
		}

		archivar.Tasks = archivar.Tasks[len(filtered):]

		if len(archivar.Tasks) > 0 {

			time.AfterFunc(1*time.Second, func() {
				processTasks(archivar)
			})

		} else {

			archivar.Busy = false

		}

	}

}

func processThreadListing(archivar *Archivar, schema schemas.ThreadListing, prev_task ArchivarTask) {

	if schema.Data.After != nil {

		var next_task ArchivarTask

		next_task.Subreddit = archivar.Subreddit
		next_task.Type = prev_task.Type
		next_task.Query = prev_task.Query
		next_task.Page  = *schema.Data.After
		next_task.Count = prev_task.Count+100

		archivar.Tasks = append(archivar.Tasks, next_task)

	}

	if len(schema.Data.Children) > 0 {

		for c := 0; c < len(schema.Data.Children); c++ {

			var child = schema.Data.Children[c]

			buffer, err := json.MarshalIndent(child, "", "\t")

			if err == nil {

				if child.Data.Identifier != "" {

					if archivar.Cache.Exists("threads/" + child.Data.Identifier + ".json") == false {
						archivar.Cache.Write("threads/" + child.Data.Identifier + ".json", buffer)
					}

					if archivar.Cache.Exists("comments/" + child.Data.Identifier + ".json") == false {

						var comments_task ArchivarTask

						comments_task.Subreddit = archivar.Subreddit
						comments_task.Type = "comments"
						comments_task.Query = ""
						comments_task.Page = child.Data.Identifier
						comments_task.Count = 0

						archivar.Scraper.DeferRequest(comments_task.ToURL(), func(buffer []byte, status int) {

							if status == 200 && len(buffer) > 0 {
								archivar.Cache.Write("comments/" + child.Data.Identifier + ".json", buffer)
							}

						})

					}

				}

			}

		}

	}

}

func NewArchivar(subreddit string, folder string) Archivar {

	var archivar Archivar

	archivar.Cache     = NewCache(folder)
	archivar.Scraper   = NewScraper(1)
	archivar.Subreddit = subreddit
	archivar.Errors = make([]string, 0)
	archivar.Tasks = make([]ArchivarTask, 0)

	return archivar

}

func (archivar *Archivar) DeferScrape(typ string, query string) {

	var task ArchivarTask

	if typ != "" && query != "" {
		task.Subreddit = archivar.Subreddit
		task.Type = typ
		task.Query = query
	} else if typ != "" {
		task.Subreddit = archivar.Subreddit
		task.Type = typ
	}


	if task.Subreddit != "" && task.Type != "" {
		archivar.Tasks = append(archivar.Tasks, task)
	}

	if archivar.Busy == false {

		archivar.Busy = true

		time.AfterFunc(1*time.Second, func() {
			processTasks(archivar)
		})

	}

}

