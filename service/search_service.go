package service

import (
	"github.com/otz1/otz/conv"
	"time"

	"github.com/otz1/otz/client"
	"github.com/otz1/otz/entity"
)

const (
	ResultsPerPage = 10
)

// SearchService ...
type SearchService struct{}

func NewSearchService() *SearchService {
	return &SearchService{}
}

func timedEvent(evt func()) time.Duration {
	startTime := time.Now()
	evt()
	return time.Now().Sub(startTime)
}

// Search ...
func (s *SearchService) Search(query string) entity.SearchResponse {
	var scraperResp *entity.ScrapeResponse
	elapsedTime := timedEvent(func() {
		scraperClient := client.NewScraperClient()
		scraperResp = scraperClient.Scrape(query)
	})

	// perhaps we could move this.
	var results []entity.SearchResult
	for _, result := range scraperResp.Results {
		results = append(results, conv.ToSearchResult(result))
	}

	numPages := max(len(results)/ResultsPerPage, 1)

	return entity.SearchResponse{
		Query:   query,
		Results: results,
		Measurements: entity.MeasurementDetail{
			ElapsedTime: elapsedTime,
			ResultCount: len(results),
		},
		NumPages: numPages,
	}
}

func buildResult() entity.SearchResult {
	return entity.SearchResult{
		Title:           "title",
		Snippet:         "this is a snippet from the webpage",
		Ranking:         1,
		ImageSource:     "http://placehold.it/256x256",
		ThumbnailSource: "http://placehold.it/256x256",
		Href:            "https://felixangell.com",
	}
}

// go only has min and max for floats...
// to avoid casting lets just delcare our own for now.
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
