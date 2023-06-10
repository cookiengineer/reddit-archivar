package structs

import "strconv"
import "strings"

type Task struct {
	Subreddit string `json:"subreddit"`
	Type      string `json:"type"`
	Query     string `json:"query"`
	Count     int    `json:"count"`
}

func NewTask(subreddit string, typ string, query string) Task {

	var task Task

	task.Subreddit = strings.TrimSpace(subreddit)
	task.Type = typ
	task.Query = query
	task.Count = 0

	return task

}

func (task *Task) ToURL(identifier string) string {

	var url string

	if task.Type == "new" || task.Type == "hot" || task.Type == "top" {

		url = "https://old.reddit.com/r/" + task.Subreddit + "/" + task.Type + ".json"
		url += "?t=all"
		url += "&limit=100"

		if task.Count != 0 {
			url += "&count=" + strconv.Itoa(task.Count)
		}

		if identifier != "" {
			url += "&after=" + identifier
		}

	} else if task.Type == "comments" {

		url = "https://old.reddit.com/r/" + task.Subreddit + "/comments/" + identifier + ".json"

	} else if task.Type == "search" {

		url = "https://old.reddit.com/r/" + task.Subreddit + "/search.json"
		url += "?q=" + task.Query
		url += "&restrict_sr=on"
		url += "&t=all"

		if task.Count != 0 {
			url += "&count=" + strconv.Itoa(task.Count)
		}

		if identifier != "" {
			url += "&after=" + identifier
		}

	}

	return url

}
