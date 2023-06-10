package schemas

type ThreadListing struct {
	Kind string `json:"kind"`
	Data struct {
		After     *string  `json:"after"`
		Dist      int      `json:"dist"`
		ModHash   string   `json:"modhash"`
		GeoFilter string   `json:"geofilter"`
		Children  []Thread `json:"children"`
		Before    *string  `json:"before"`
	} `json:"data"`
}

type Thread struct {
	Kind string `json:"kind"`
	Data struct {
		Identifier            string          `json:"id"`
		Name                  string          `json:"name"`
		ApprovedAtUTC         *string         `json:"approved_at_utc"`
		BannedAtUTC           *string         `json:"banned_at_utc"`
		ApprovedBy            *string         `json:"approved_by"`
		Subreddit             string          `json:"subreddit"`
		SubredditIdentifier   string          `json:"subreddit_id"`
		SubredditNamePrefixed string          `json:"subreddit_name_prefixed"`
		Title                 string          `json:"title"`
		Permalink             string          `json:"permalink"`
		Domain                string          `json:"domain"`
		URL                   string          `json:"url"`
		Selftext              string          `json:"selftext"`
		SelftextHTML          string          `json:"selftext_html"`
		AuthorFullname        string          `json:"author_fullname"`
		Hidden                bool            `json:"hidden"`
		IsOriginalContent     bool            `json:"is_original_content"`
		Created               Timestamp       `json:"created"`
		Replies               *CommentListing `json:"replies"`
	} `json:"data"`
}
