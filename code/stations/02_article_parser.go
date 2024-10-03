package stations

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"tobloggan/code/contracts"
)

type ArticleParser struct{}

// TODO: given a contracts.SourceFile, parse the JSON metadata and save the body on a contracts.Article.
// input: contracts.SourceFile
// output: contracts.Article

func (this *ArticleParser) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.SourceFile:
		var article contracts.Article
		myString := string(input)

		// convert the bytes to string
		parts := strings.Split(myString, "+++")

		// json unmarshall to an contracts.Article
		err := json.Unmarshal([]byte(parts[0]), &article)
		if err != nil {
			error := fmt.Errorf("Error unmarchsaling JSON: %w : %w", err, errMalformedContent)
			output(error)
			return // todo: needed?
		}

		// Assign body of input to body of contracts.Article
		article.Body = parts[1]

		// Output the created contracts.Article
		output(article)
	default:
		output(input)
	}
}

var errMalformedContent = errors.New("malformed content")
