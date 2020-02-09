package service

import (
	"time"

	"github.com/otz1/otz/client"
	"github.com/otz1/otz/entity"
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

	results := []entity.SearchResult{
		entity.SearchResult{
			Title:           "title",
			Snippet:         "this is a snippet from the webpage",
			Ranking:         1,
			ImageSource:     "http://placehold.it/256x256",
			ThumbnailSource: "http://placehold.it/256x256",
			Href:            "https://felixangell.com",
		},
	}

	elapsedTime := time.Now().Sub(startTime)

	return entity.SearchResponse{
		Query:   query,
		Results: results,
		Measurements: entity.MeasurementDetail{
			ElapsedTime: elapsedTime,
			ResultCount: len(results),
		},
		NumPages: 123,
	}
}
