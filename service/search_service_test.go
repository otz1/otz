package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// in the future this will be less flaky when the
// results are cached properly.
func TestItConvertsSearchesFromScraper(t *testing.T) {
	query := "how to make pancakes"

	client := NewSearchService()
	clientResponse := client.Search(query)

	ss := NewSearchService()
	ssResponse := ss.Search(query)

	assert.Equal(t, clientResponse.Results[0].Title, ssResponse.Results[0].Title)
	assert.Equal(t, clientResponse.Results[0].Href, ssResponse.Results[0].Href)
	assert.Equal(t, len(clientResponse.Results), len(ssResponse.Results))
}

func TestSearchReturnsMeasurementsWithElapsedTime(t *testing.T) {
	// given a search service
	ss := NewSearchService()

	// when we invoke the service
	resp := ss.Search("hello, world")

	// then we have a non zero elapsed time
	assert.NotZero(t, resp.Measurements.ElapsedTime)
}

func TestResultsPerPage(t *testing.T) {
	// given a search service
	ss := NewSearchService()

	// when we invoke the service
	resp := ss.Search("hello, world")

	numPages := max(len(resp.Results)/ResultsPerPage, 1)
	assert.Equal(t, numPages, resp.NumPages)
}
