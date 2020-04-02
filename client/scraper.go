package client

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/otz1/otz/entity"
	"github.com/parnurzeal/gorequest"
)

type ScraperClient struct {}

func (s *ScraperClient) Scrape(query string) *entity.ScrapeResponse {
	scrapeRequest := entity.ScrapeRequest {
		Query: query,
		Source: "DDG",
	}

	_, body, errs := gorequest.New().Post("https://otzs.otzaf.org/scrape").Send(scrapeRequest).End()
	if len(errs) > 0 {
		panic(errs)
	}

	resp := &entity.ScrapeResponse{}
	if err := jsoniter.Unmarshal([]byte(body), resp); err != nil {
		panic(err)
	}
	return resp
}

func NewScraperClient() *ScraperClient {
	return &ScraperClient{}
}
