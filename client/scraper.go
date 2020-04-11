package client

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/otz1/otz/entity"
	"github.com/otz1/otz/util"
	"github.com/parnurzeal/gorequest"
)

var scrapeEndpoint = map[string]string{
	"prod":    "https://otzs.otzaf.org/scrape",
	"staging": "http://otzs.qa.otzaf.org/scrape",
	// local speaks to QA.
	"local": "http://otzs.qa.otzaf.org/scrape",
}

type ScraperClient struct{}

func (s *ScraperClient) getEndpoint() string {
	env := util.GetEnv("ENVIRONMENT", "local")
	endpoint, ok := scrapeEndpoint[env]
	if !ok {
		sentry.CaptureException(fmt.Errorf("not such environment '%s'", env))
	}
	return endpoint
}

func (s *ScraperClient) Scrape(query string) *entity.ScrapeResponse {
	scrapeRequest := entity.ScrapeRequest{
		Query: query,
	}

	_, body, errs := gorequest.
		New().
		Post(s.getEndpoint()).
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
