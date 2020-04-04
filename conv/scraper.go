package conv

import (
	"fmt"
	"github.com/otz1/otz/entity"
	"strings"
)

// rule of thumb is that we name
// file after the domain we are converting
// FROM, in this case the domain we
// are converting from is the scraper client.

func EmphasizeSnippetSearchTerms(searchTerms []string, sr entity.SearchResult) entity.SearchResult {

	// TODO we could pre-allocate here because it will always
	// be searchTerms * 2.

	var replacementSet []string
	for _, st := range searchTerms {
		if len(st) < 2 {
			continue
		}
		replacedWord := fmt.Sprintf("<span className='keyword'>%s</span>", st)
		replacementSet = append(replacementSet, []string{
			st, replacedWord,
		}...)
	}

	replacer := strings.NewReplacer(replacementSet...)
	emphasizedSnippet := replacer.Replace(sr.Snippet)

	return entity.SearchResult{
		Title:           sr.Title,
		Snippet:         emphasizedSnippet,
		Ranking:         sr.Ranking,
		ImageSource:     sr.ImageSource,
		ThumbnailSource: sr.ThumbnailSource,
		Href:            sr.Href,
	}
}

func ToSearchResult(scrapedResult entity.ScrapeResult) entity.SearchResult {
	// TODO finish the conversion!
	return entity.SearchResult{
		Title: scrapedResult.Title,
		Href: scrapedResult.Href,
		Snippet: scrapedResult.Snippet,
	}
}