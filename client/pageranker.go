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
	"prod":    "https://otzpr.otzaf.org/fetch",
	"staging": "https://otzpr.qa.otzaf.org/fetch",
	// local speaks to QA.
	"local": "https://otzpr.qa.otzaf.org/fetch",
}

type PageRankerClient struct{}

func (s *PageRankerClient) getEndpoint() string {
	env := util.GetEnv("ENVIRONMENT", "local")
	endpoint, ok := scrapeEndpoint[env]
	if !ok {
		sentry.CaptureException(fmt.Errorf("not such environment '%s'", env))
	}
	return endpoint
}

func (s *PageRankerClient) Fetch(query string) *entity.PageRankerResponse {
	fetchRequest := entity.PageRankerRequest{
		Query: query,
	}

	_, body, errs := gorequest.
		New().
		Post(s.getEndpoint()).
		AppendHeader("Content-Type", "application/json").
		AppendHeader("Accept", "application/json").
		AppendHeader("SITE-CODE", "OTZIT_UK"). // TODO deduce this rather than hardcode
		Send(fetchRequest).
		End()

	if len(errs) > 0 {
		for _, err := range errs {
			sentry.CaptureException(err)
		}
		panic(errs)
	}

	resp := &entity.PageRankerResponse{}
	if err := jsoniter.Unmarshal([]byte(body), resp); err != nil {
		sentry.CaptureException(fmt.Errorf("failed to unmarshal request to '%s' err '%v'", s.getEndpoint(), err))
		panic(err)
	}
	return resp
}

func NewPageRankerClient() *PageRankerClient {
	return &PageRankerClient{}
}
