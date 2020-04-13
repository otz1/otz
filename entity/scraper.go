package entity

type PageRankerRequest struct {
	Query string `json:"query"`
}

type RankedResultData struct {
	Title   string `json:"title"`
	Href    string `json:"href"`
	Snippet string `json:"snippet"`
}

type RankedSearchResult struct {
	Result RankedResultData `json:"result"`
	Score  int              `json:"score"`
}

type PageRankerResponse struct {
	Results []RankedSearchResult `json:"results"`
}
