package conv

import (
	"github.com/otz1/otz/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItEmphasizesWordsProperly(t *testing.T) {
	srWithEmphasizedSnippet := EmphasizeSnippetSearchTerms([]string{"how", "to", "make", "pancakes"}, entity.SearchResult{
		Snippet: "making pancakes is really simple how you make them to make them fast",
	})

	expected := "making <span className='keyword'>pancakes</span> is really simple <span className='keyword'>how</span> you <span className='keyword'>make</span> them <span className='keyword'>to</span> <span className='keyword'>make</span> them fast"
	assert.Equal(t, expected, srWithEmphasizedSnippet.Snippet)
}

func TestItDoesntEmphasizeWordsUnderTwoChars(t *testing.T) {
	srWithEmphasizedSnippet := EmphasizeSnippetSearchTerms([]string{"this", "is", "a", "test", "to"}, entity.SearchResult{
		Snippet: "a test this is o test u snippet word blah toast test",
	})

	expected := "a <span className='keyword'>test</span> <span className='keyword'>this</span> <span className='keyword'>is</span> o <span className='keyword'>test</span> u snippet word blah <span className='keyword'>to</span>ast <span className='keyword'>test</span>"
	assert.Equal(t, expected, srWithEmphasizedSnippet.Snippet)
}