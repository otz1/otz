package entity

type ScrapeRequest struct {
	Query string `json:"query"`
	Source string `json:"source"`
}

type ScrapeResult struct {
	Title string
	Href string
}

type ScrapeResponse struct {
	OriginalQuery string
	Results []ScrapeResult
}
