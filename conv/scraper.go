package conv

import "github.com/otz1/otz/entity"

// rule of thumb is that we name
// file after the domain we are converting
// FROM, in this case the domain we
// are converting from is the scraper client.

func ToSearchResult(scrapedResult entity.ScrapeResult) entity.SearchResult {
	// TODO finish the conversion!
	return entity.SearchResult{
		Title: scrapedResult.Title,
		Href: scrapedResult.Href,
		Snippet: scrapedResult.Snippet,
	}
}