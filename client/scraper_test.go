package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItReturnsScrapeResponse(t *testing.T) {
	client := NewScraperClient()
	response := client.Scrape("how to make pancakes")
	assert.NotNil(t, response)
}