package service

import (
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

// Search ...
func (s *SearchService) Search(query string) entity.SearchResponse {
	startTime := time.Now()

	pr := client.NewPageRankerClient()

	// check cache.
	// fail -> check db for enough keywords. store them in db
	// 			store them in cache.

	// succ -> page rank

	pr.GetRanking(query)

	var results []entity.SearchResult
	for i := 0; i < 40; i++ {
		results = append(results, buildResult())
	}

	elapsedTime := time.Now().Sub(startTime)
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
