package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/otz1/otz/api"
)

type SearchResult struct {
	Href string `json:"href"`
}

type MeasurementDetail struct {
	ElapsedTime time.Duration `json:"elapsedTime"`
	ResultCount uint64        `json:"resultCount"`
}

type SearchResponse struct {
	Query        string            `json:"query"`
	Results      []SearchResult    `json:"results"`
	Measurements MeasurementDetail `json:"measurements"`
}

func (ctx *Controller) Search(c *gin.Context) {
	query := c.Query("query")

	log.Println("Search request for", query)

	api.ProcessSearch(query)

	results := []SearchResult{
		SearchResult{"http://google.com/test"},
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
