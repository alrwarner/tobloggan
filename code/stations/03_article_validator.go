package stations

import (
	"tobloggan/code/contracts"
)

type ArticleValidator struct{}

func (this *ArticleValidator) Do(input any, output func(any)) {
	//    TODO: given a contracts.Article, validate the Slug and the Title fields and emit the contracts.Article (or an error)
	//    input: contracts.Article
	//    output: contracts.Article (or error)

	switch input := input.(type) {
	case contracts.Article:

		// Check the Title of the contracts.Article
		if len(input.Title) == 0 {
			output(errMalformedContent)
		} else if len(input.Title) > 150 {
			output(errMalformedContent)
		}

		// check slug characters are valid
		if !isValidSlug(input.Slug, validSlugCharacters) {
			output(errMalformedContent)
		}

		// TODO: Check date range

		// Output article, there were no problems
		output(input)
	default:
		output(input)
	}
}

func newRuneSet(chars string) map[rune]struct{} {
	set := make(map[rune]struct{})
	for _, character := range chars { // loop over the characters provided that are valid
		set[character] = struct{}{} // store each valid character into a map
	}
	return set
}

func isValidSlug(slug string, validSlugCharacters map[rune]struct{}) bool {
	for _, character := range slug { // range over all the characters in the slug
		if _, exists := validSlugCharacters[character]; !exists {
			return false // Character not found in valid characters
		}
	}
	return true
}

var validSlugCharacters = newRuneSet("abcdefghijklmnopqrstuvwxyz0123456789-/")
