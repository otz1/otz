package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	jsoniter "github.com/json-iterator/go"
)

type GoogleResult struct {
	Title           string `json:"title"`
	Snippet         string `json:"snippet"`
	Ranking         int    `json:"ranking"`
	ImageSource     string `json:"image_source"`
	ThumbnailSource string `json:"thumbnail_source"`
}

// SearchRequest will use Googles API for search results
// to search for the given query.
func SearchRequest(query string) []GoogleResult {
	key := os.Getenv("GOOGLE_KEY")
	cx := os.Getenv("GOOGLE_CX")

	query = url.PathEscape(query)

	reqURL := fmt.Sprintf("https://www.googleapis.com/customsearch/v1/siterestrict?key=%s&cx=%s&q=%s", key, cx, query)
	resp, err := http.Get(reqURL)
	if err != nil {
		log.Println(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	type resultSet struct {
		Kind string `json:"kind"`
		URL  struct {
			Type     string `json:"type"`
			Template string `json:"template"`
		} `json:"url"`
		Items []struct {
			Title       string `json:"title"`
			HTMLTitle   string `json:"htmlTitle"`
			Snippet     string `json:"snippet"`
			HTMLSnippet string `json:"htmlSnippet"`
			PageMap     struct {
				Thumbnail struct {
					Width  int    `json:"width"`
					Height int    `json:"height"`
					Source string `json:"src"`
				} `json:"thumbnail"`
				Image struct {
					Source string `json:"src"`
				} `json:"image"`
			} `json:"pagemap"`
		} `json:"items"`
	}
	var results resultSet
	if err := jsoniter.Unmarshal(data, &results); err != nil {
		log.Println(err)
	}

	items := []GoogleResult{}
	for idx, item := range results.Items {
		items = append(items, GoogleResult{
			Title:           item.Title,
			Snippet:         item.Snippet,
			Ranking:         idx,
			ImageSource:     item.PageMap.Image.Source,
			ThumbnailSource: item.PageMap.Thumbnail.Source,
		})
	}
	return items
}
