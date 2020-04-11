package client

import (
	"github.com/getsentry/sentry-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/otz1/otz/entity"
	"github.com/parnurzeal/gorequest"
)

type ScraperClient struct{}

func (s *ScraperClient) Scrape(query string) *entity.ScrapeResponse {
	scrapeRequest := entity.ScrapeRequest{
		Query:  query,
		Source: "DDG",
	}

	_, body, errs := gorequest.
		New().
		Post("https://otzs.otzaf.org/scrape"). // TODO get from env variables.
		AppendHeader("Content-Type", "application/json").
		AppendHeader("Accept", "application/json").
		AppendHeader("SITE-CODE", "OTZIT_UK"). // TODO deduce this rather than hardcode
		Send(scrapeRequest).
		End()

	if len(errs) > 0 {
		for _, err := range errs {
			sentry.CaptureException(err)
		}
		panic(errs)
	}

	resp := &entity.ScrapeResponse{}
	if err := jsoniter.Unmarshal([]byte(body), resp); err != nil {
		sentry.CaptureException(err)
		panic(err)
	}
	return resp
}

func NewScraperClient() *ScraperClient {
	return &ScraperClient{}
}
