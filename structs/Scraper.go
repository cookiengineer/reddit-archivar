package structs

import "reddit-archivar/console"
import "reddit-archivar/utils"
import "net/http"
import "io/ioutil"
import "strconv"
import "strings"
import "time"

var CONTENT_TYPES []string = []string{
	"application/gzip",
	"application/json",
	"application/ld+json",
	"application/octet-stream",
	"application/rss+xml",
	"application/x-bzip2",
	"application/x-gzip",
	"application/xml",
	"application/zip",
	"text/html",
	"text/plain",
	"text/xml",
}

type Callback func([]byte, int)

type ScraperTask struct {
	Url      string
	Callback Callback
}

type Scraper struct {
	Busy      bool
	Limit     int
	Tasks     []ScraperTask
	Headers   map[string]string
	Throttled bool
}

func processRequests(scraper *Scraper) {

	var filtered []ScraperTask
	var limit int = scraper.Limit

	if scraper.Throttled == true {
		limit = 1
	}

	for t := 0; t < len(scraper.Tasks); t++ {

		if len(filtered) < limit {
			filtered = append(filtered, scraper.Tasks[t])
		} else {
			break
		}

	}

	if len(filtered) > 0 {

		for f := 0; f < len(filtered); f++ {

			var task = filtered[f]

			buffer, status := scraper.Request(task.Url)

			task.Callback(buffer, status)

		}

		scraper.Tasks = scraper.Tasks[len(filtered):]

		if len(scraper.Tasks) > 0 {

			time.AfterFunc(1*time.Second, func() {
				processRequests(scraper)
			})

		} else {

			scraper.Busy = false

		}

	}

}

func IsScraper(scraper Scraper) bool {

	var result bool = true

	return result

}

func NewScraper(limit int) Scraper {

	if limit <= 0 {
		limit = 1
	}

	var scraper Scraper

	scraper.Busy = false
	scraper.Limit = limit
	scraper.Tasks = make([]ScraperTask, 0)
	scraper.Headers = make(map[string]string, 0)
	scraper.Throttled = false

	scraper.Headers["Accept"] = "application/json"
	scraper.Headers["Accept-Encoding"] = "gzip"
	scraper.Headers["Accept-Language"] = "en-US,en;q=0.5"
	scraper.Headers["Cache-Control"] = "no-cache"
	scraper.Headers["Connection"] = "keep-alive"
	scraper.Headers["Host"] = "old.reddit.com"
	scraper.Headers["Pragma"] = "no-cache"
	scraper.Headers["Referer"] = "https://old.reddit.com/"
	scraper.Headers["Sec-Fetch-Dest"] = "document"
	scraper.Headers["Sec-Fetch-Mode"] = "navigate"
	scraper.Headers["Sec-Fetch-Site"] = "none"
	scraper.Headers["Sec-Fetch-User"] = "?1"
	scraper.Headers["TE"] = "trailers"
	scraper.Headers["Upgrade-Insecure-Requests"] = "1"
	scraper.Headers["User-Agent"] = "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/113.0"

	return scraper

}

func (scraper *Scraper) DeferRequest(url string, callback Callback) {

	scraper.Tasks = append(scraper.Tasks, ScraperTask{
		Url:      url,
		Callback: callback,
	})

	if scraper.Busy == false {

		scraper.Busy = true

		time.AfterFunc(1*time.Second, func() {
			processRequests(scraper)
		})

	}

}

func (scraper *Scraper) Request(url string) ([]byte, int) {

	var buffer []byte
	var content_encoding string
	var content_type string
	var status_code int

	client := &http.Client{}
	client.CloseIdleConnections()

	request, err1 := http.NewRequest("GET", url, nil)

	if err1 == nil {

		for key, val := range scraper.Headers {
			request.Header.Set(key, val)
		}

		response, err2 := client.Do(request)

		if err2 == nil {

			status_code = response.StatusCode

			if status_code == 200 || status_code == 304 {

				if len(response.Header["Content-Type"]) > 0 {
					content_type = response.Header["Content-Type"][0]
				}

				if len(response.Header["Content-Encoding"]) > 0 {
					content_encoding = response.Header["Content-Encoding"][0]
				}

				var valid bool = false

				for c := 0; c < len(CONTENT_TYPES); c++ {

					if strings.Contains(content_type, CONTENT_TYPES[c]) {
						valid = true
						break
					}

				}

				if valid == true {

					data, err2 := ioutil.ReadAll(response.Body)

					if err2 == nil {
						buffer = data
					}

				}

			}

		}

	}

	if content_encoding == "identitiy" {

		// Do Nothing

	} else if content_encoding == "gzip" {

		decompressed := utils.GUnzip(buffer)

		if len(decompressed) > 0 {
			buffer = decompressed
		} else {
			buffer = []byte{}
		}

	}

	if len(buffer) > 0 {

		console.Log("Request \"" + url + "\"")

	} else {

		console.Error("Request \"" + url + "\"")

		if content_type != "" {
			console.Error("Unsupported Content-Type \"" + content_type + "\"")
		}

		if status_code != 0 {
			console.Error("Unsupported Status Code \"" + strconv.Itoa(status_code) + "\"")
		}

	}

	return buffer, status_code

}

