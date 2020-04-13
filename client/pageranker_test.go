package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItReturnsScrapeResponse(t *testing.T) {
	client := NewPageRankerClient()
	response := client.Fetch("how to make pancakes")
	assert.NotNil(t, response)
}
