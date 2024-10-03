package stations

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"tobloggan/code/contracts"
)

type ArticleParser struct{}

func NewArticleParser() contracts.Station {
	return &ArticleParser{}
}

// TODO: given a contracts.SourceFile, parse the JSON metadata and save the body on a contracts.Article.
// input: contracts.SourceFile
// output: contracts.Article

func (this *ArticleParser) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.SourceFile:
		var article contracts.Article
		myString := string(input)

		if !strings.Contains(myString, "+++") {
			error := fmt.Errorf("Error article doesn't contain seperator +++: %w", errMalformedContent)
			output(error)
			return
		}

		// convert the bytes to string
		parts := strings.Split(myString, "+++")

		// json unmarshall to an contracts.Article
		err := json.Unmarshal([]byte(parts[0]), &article)
		if err != nil {
			error := fmt.Errorf("Error unmarchsaling JSON: %w : %w", err, errMalformedContent)
			output(error)
			return
		}

		// Assign body of input to body of contracts.Article

		article.Body = strings.ReplaceAll(parts[1], "\n", "")

		// Output the created contracts.Article
		output(article)
		return
	default:
		output(input)
		return
	}
}

var (
	errMalformedContent = errors.New("malformed content")
)
