package stations

import (
	"errors"

	"tobloggan/code/contracts"
	"tobloggan/code/markdown"
)

type Markdown interface {
	Convert(content string) (string, error)
}

type MarkdownConverter struct {
	converter Markdown
}

func NewMarkdownConverter(converter Markdown) *MarkdownConverter {
	return &MarkdownConverter{converter: converter}
}

func (this MarkdownConverter) Do(input any, output func(any)) {
	var err error
	converter := markdown.NewConverter()
	switch input := input.(type) {
	case contracts.Article:
		input.Body, err = converter.Convert(input.Body)
		if err != nil {
			output(errors.New("problem converting article to markdown: " + err.Error()))
		} else {
			output(input)
		}
	default:
		output(input)
	}
}
