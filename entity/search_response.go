package entity

// SearchResponse ...
type SearchResponse struct {
	Query        string            `json:"query"`
	Results      []SearchResult    `json:"results"`
	Measurements MeasurementDetail `json:"measurements"`
	NumPages     int               `json:"numPages"`
	SearchTerms  []string          `json:"searchTerms"`
}
