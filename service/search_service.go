package service

import (
	"github.com/otz1/otz/conv"
	"strings"
	"time"
	"unicode"

	"github.com/otz1/otz/client"
	"github.com/otz1/otz/entity"
)

const (
	ResultsPerPage = 10
)

// SearchService ...
type SearchService struct {
	scraperClient *client.PageRankerClient
}

func NewSearchService() *SearchService {
	return &SearchService{
		scraperClient: client.NewPageRankerClient(),
	}
}

func timedEvent(evt func()) time.Duration {
	startTime := time.Now()
	evt()
	return time.Now().Sub(startTime) / time.Millisecond
}

func extractSearchTerms(query string) []string {
	// naive algorithm, replace each non word character
	// with a space, then we split the string by spaces.
	// this allows the support of other languages rather than
	// using a complex regular expr.
	runes := []rune(query)
	for i, ch := range runes {
		if unicode.IsSymbol(ch) || unicode.IsPunct(ch) {
			runes[i] = ' '
		}
	}
	return strings.Fields(string(runes))
}

// Search ...
func (s *SearchService) Search(query string) entity.SearchResponse {
	var fetchResp *entity.PageRankerResponse
	elapsedTime := timedEvent(func() {
		fetchResp = s.scraperClient.Fetch(query)
	})

	searchTerms := extractSearchTerms(query)

	// perhaps we could move this into a converter
	results := make([]entity.SearchResult, len(fetchResp.Results))
	for i, result := range fetchResp.Results {
		converted := conv.ToSearchResult(result)
		emphasized := conv.EmphasizeSnippetSearchTerms(searchTerms, converted)
		results[i] = emphasized
	}

	numPages := max(len(results)/ResultsPerPage, 1)

	return entity.SearchResponse{
		Query:   query,
		Results: results,
		Measurements: entity.MeasurementDetail{
			ElapsedTime: elapsedTime,
			ResultCount: len(results),
		},
		NumPages:    numPages,
		SearchTerms: searchTerms,
	}
}

// go only has min and max for floats...
// to avoid casting lets just declare our own for now.
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
