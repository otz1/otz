package entity

type ScrapeRequest struct {
	Query string `json:"query"`
	Source string `json:"source"`
}

type ScrapeResult struct {
	Title string
	Href string
	Snippet string
}

type ScrapeResponse struct {
	OriginalQuery string
	Results []ScrapeResult
}
