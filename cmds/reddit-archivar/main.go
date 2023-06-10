package main

import "reddit-archivar/console"
import "reddit-archivar/structs"
import "reddit-archivar/schemas"
import "encoding/json"
import "os"
import "strconv"
import "strings"

var CACHE structs.Cache
var ERRORS int
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
	ERRORS = 0
	SCRAPER = structs.NewScraper(1)
}

func scrape(task *structs.Task) {

	var url = task.ToURL("")

	if url != "" {

		SCRAPER.DeferRequest(url, func(buffer []byte) {

			var schema schemas.Listing

			err := json.Unmarshal(buffer, &schema)

			if err == nil {
				processListing(task, &schema)
			}

		})

	}

}

func processListing(task *structs.Task, schema *schemas.Listing) {

	if schema.Data.After != nil {

		task.Count += 100

		var url = task.ToURL(*schema.Data.After)

		if url != "" {

			SCRAPER.DeferRequest(url, func(buffer []byte) {

				var schema schemas.Listing

				err := json.Unmarshal(buffer, &schema)

				if err == nil {

					processListing(task, &schema)

				} else {

					ERRORS++

					os.WriteFile("/tmp/reddit-error-" + strconv.Itoa(ERRORS) + ".json", buffer, 0666)
					console.Warn(err.Error())

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

					if CACHE.Exists("listing/" + child.Data.Identifier + ".json") == false {
						CACHE.Write("listing/" + child.Data.Identifier + ".json", buffer)
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

			scrape(&task)

		}

		var task_hot = structs.NewTask(subreddit, "hot", "")
		var task_new = structs.NewTask(subreddit, "new", "")
		var task_top = structs.NewTask(subreddit, "top", "")

		scrape(&task_hot)
		scrape(&task_new)
		scrape(&task_top)

		for SCRAPER.Busy == true {
			// Wait for Scraper
		}

	} else {

		console.Error("Please provide a Subreddit as the first parameter.")
		console.Error("Example: reddit-archivar /r/Malware")
		os.Exit(1)

	}

}
