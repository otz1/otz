package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItReturnsResultsFromGoogle(t *testing.T) {
	results := SearchRequest("how to make pancakes")
	assert.NotEmpty(t, results[0].Title)
	assert.Equal(t, 10, len(results))
}
