package main

import "reddit-archivar/console"
import "reddit-archivar/structs"
import "encoding/json"
import "os"
import "strings"

var ARCHIVAR structs.Archivar
var KEYWORDS []string

func init() {

	cwd, err1 := os.Getwd()

	if err1 == nil {

		stat, err2 := os.Stat(cwd + "/keywords.json")

		if err2 == nil && stat.IsDir() == false {

			buffer, err3 := os.ReadFile(cwd + "/keywords.json")

			if err3 == nil {
				json.Unmarshal(buffer, &KEYWORDS)
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

			if strings.HasSuffix(cwd, "/build") {
				cwd = cwd[0:len(cwd)-6]
			}

			ARCHIVAR = structs.NewArchivar(subreddit, cwd + "/archive/" + subreddit)

		} else {

			ARCHIVAR = structs.NewArchivar(subreddit, "/tmp/reddit-archivar/" + subreddit)

		}

		ARCHIVAR.DeferScrape("hot", "")
		ARCHIVAR.DeferScrape("new", "")
		ARCHIVAR.DeferScrape("top", "")

		for k := 0; k < len(KEYWORDS); k++ {
			ARCHIVAR.DeferScrape("search", KEYWORDS[k])
		}

		for ARCHIVAR.Busy == true {
			// Wait for Scraper
		}

	} else {

		console.Error("Please provide a Subreddit as the first parameter.")
		console.Error("Example: reddit-archivar /r/Malware")
		os.Exit(1)

	}

}
