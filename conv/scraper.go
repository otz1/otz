package conv

import (
	"fmt"
	"strings"

	"github.com/otz1/otz/entity"
)

// rule of thumb is that we name
// file after the domain we are converting
// FROM, in this case the domain we
// are converting from is the scraper client.

func EmphasizeSnippetSearchTerms(searchTerms []string, sr entity.SearchResult) entity.SearchResult {
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

func ToSearchResult(rankedSearchResult entity.RankedSearchResult) entity.SearchResult {
	resultData := rankedSearchResult.Result
	return entity.SearchResult{
		Title:   resultData.Title,
		Href:    resultData.Href,
		Snippet: resultData.Snippet,
	}
}
