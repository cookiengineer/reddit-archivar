package main

import "reddit-archivar/console"
import "reddit-archivar/structs"
import "reddit-archivar/schemas"
import "encoding/json"
import "os"
import "strings"

var CACHE structs.Cache
var SCRAPER structs.Scraper
var KEYWORDS []string = []string{
	"CVE",
	"RCE",
	"vulnerability",
	"exploit",
	"zeroday",
	"0-day",
	"ransomware",
	"breach",
	"leak",
}

func init() {
	SCRAPER = structs.NewScraper(1)
}

func scrapeThreads(task *structs.Task) {

	var url = task.ToURL("")

	if url != "" {

		SCRAPER.DeferRequest(url, func(buffer []byte) {

			var schema schemas.ThreadListing

			err := json.Unmarshal(buffer, &schema)

			if err == nil {
				processThreads(task, &schema)
			}

		})

	}

}

func processThreads(task *structs.Task, schema *schemas.ThreadListing) {

	if schema.Data.After != nil {

		task.Count += 100

		var url = task.ToURL(*schema.Data.After)

		if url != "" {

			SCRAPER.DeferRequest(url, func(buffer []byte) {

				var schema schemas.ThreadListing

				err := json.Unmarshal(buffer, &schema)

				if err == nil {
					processThreads(task, &schema)
				}

			})

		}

	}

	if len(schema.Data.Children) > 0 {

		for c := 0; c < len(schema.Data.Children); c++ {

			child := schema.Data.Children[c]
			buffer, err := json.MarshalIndent(child, "", "\t")

			if err == nil {

				if child.Data.Identifier != "" {

					if CACHE.Exists("threads/" + child.Data.Identifier + ".json") == false {
						CACHE.Write("threads/" + child.Data.Identifier + ".json", buffer)
					}

					if CACHE.Exists("comments/" + child.Data.Identifier + ".json") == false {

						comments_task := structs.NewTask(task.Subreddit, "comments", "")
						comments_url := comments_task.ToURL(child.Data.Identifier)

						if comments_url != "" {

							buffer := SCRAPER.Request(comments_url)

							if len(buffer) > 0 {
								CACHE.Write("comments/" + child.Data.Identifier + ".json", buffer)
							}

						}

					}

				}

			}

		}

	}

}

func main() {

	var subreddit string

	if len(os.Args) == 2 {

		if strings.HasPrefix(os.Args[1], "/r/") {

			tmp := os.Args[1][3:]

			if strings.Contains(tmp, "?") {
				tmp = tmp[0:strings.Index(tmp, "?")]
			}

			if strings.Contains(tmp, "/") {
				subreddit = strings.TrimSpace(tmp[0:strings.Index(tmp, "/")])
			} else {
				subreddit = strings.TrimSpace(tmp)
			}

		} else if strings.HasPrefix(os.Args[1], "https://reddit.com/r/") {

			tmp := os.Args[1][21:]

			if strings.Contains(tmp, "?") {
				tmp = tmp[0:strings.Index(tmp, "?")]
			}

			if strings.Contains(tmp, "/") {
				subreddit = strings.TrimSpace(tmp[0:strings.Index(tmp, "/")])
			} else {
				subreddit = strings.TrimSpace(tmp)
			}

		} else if strings.HasPrefix(os.Args[1], "https://old.reddit.com/r/") {

			tmp := os.Args[1][25:]

			if strings.Contains(tmp, "?") {
				tmp = tmp[0:strings.Index(tmp, "?")]
			}

			if strings.Contains(tmp, "/") {
				subreddit = strings.TrimSpace(tmp[0:strings.Index(tmp, "/")])
			} else {
				subreddit = strings.TrimSpace(tmp)
			}

		} else if strings.HasPrefix(os.Args[1], "https://www.reddit.com/r/") {

			tmp := os.Args[1][25:]

			if strings.Contains(tmp, "?") {
				tmp = tmp[0:strings.Index(tmp, "?")]
			}

			if strings.Contains(tmp, "/") {
				subreddit = strings.TrimSpace(tmp[0:strings.Index(tmp, "/")])
			} else {
				subreddit = strings.TrimSpace(tmp)
			}

		}

	}

	console.Group("reddit-archivar: Command-Line Arguments")
	console.Inspect(struct{ Subreddit string }{Subreddit: subreddit})
	console.GroupEnd("")

	if subreddit != "" {

		cwd, err := os.Getwd()

		if err == nil {
			CACHE = structs.NewCache(cwd + "/archive/" + subreddit)
		} else {
			CACHE = structs.NewCache("/tmp/reddit-archivar/" + subreddit)
		}

		for k := 0; k < len(KEYWORDS); k++ {

			var task = structs.NewTask(subreddit, "search", KEYWORDS[k])

			scrapeThreads(&task)

		}

		var task_hot = structs.NewTask(subreddit, "hot", "")
		var task_new = structs.NewTask(subreddit, "new", "")
		var task_top = structs.NewTask(subreddit, "top", "")

		scrapeThreads(&task_hot)
		scrapeThreads(&task_new)
		scrapeThreads(&task_top)

		for SCRAPER.Busy == true {
			// Wait for Scraper
		}

	} else {

		console.Error("Please provide a Subreddit as the first parameter.")
		console.Error("Example: reddit-archivar /r/Malware")
		os.Exit(1)

	}

}
