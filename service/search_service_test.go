package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchReturnsMeasurementsWithElapsedTime(t *testing.T) {
	// given a search service
	ss := NewSearchService()

	// when we invoke the service
	resp := ss.Search("hello, world")

	// then we have a non zero elapsed time
	assert.NotZero(t, resp.Measurements.ElapsedTime)
}
