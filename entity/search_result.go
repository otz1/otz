package entity

// SearchResult is an individual result
// that is returned.
type SearchResult struct {
	Title           string `json:"title"`
	Snippet         string `json:"snippet"`
	Ranking         int    `json:"ranking"`
	ImageSource     string `json:"imageSource"`
	ThumbnailSource string `json:"thumbnailSource"`
	Href            string `json:"href"`
}