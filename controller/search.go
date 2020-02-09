package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/otz1/otz/api"
)

// SearchResult is an individual result
// that is returned.
type SearchResult struct {
	Title           string `json:"title"`
	Snippet         string `json:"snippet"`
	Ranking         int    `json:"ranking"`
	ImageSource     string `json:"imageSource"`
	ThumbnailSource string `json:"thumbnailSource"`
	Href            string `json:"href"`
}

// MeasurementDetail is information that provides
// how long the request took and how many results were given
type MeasurementDetail struct {
	// perhaps elapsed time could be a float?
	ElapsedTime time.Duration `json:"elapsedTime"`
	ResultCount uint64        `json:"resultCount"`
}

// SearchResponse ...
type SearchResponse struct {
	Query        string            `json:"query"`
	Results      []SearchResult    `json:"results"`
	Measurements MeasurementDetail `json:"measurements"`
	NumPages     int               `json:"numPages"`
}

// Search ...
func (ctx *Controller) Search(c *gin.Context) {
	query := c.Query("query")

	log.Println("Search request for", query)

	api.ProcessSearch(query)

	results := []SearchResult{
		SearchResult{
			Title:           "title",
			Snippet:         "this is a snippet from the webpage",
			Ranking:         1,
			ImageSource:     "http://placehold.it/256x256",
			ThumbnailSource: "http://placehold.it/256x256",
			Href:            "https://felixangell.com",
		},
	}
	c.JSON(http.StatusOK, SearchResponse{
		Query:   query,
		Results: results,
		Measurements: MeasurementDetail{
			ElapsedTime: time.Millisecond * 25,
			ResultCount: 2312129,
		},
	})
}
