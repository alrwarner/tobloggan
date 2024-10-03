package stations

import (
	"strings"

	"tobloggan/code/contracts"
)

type ArticleRenderer struct {
	template string
}

func (this *ArticleRenderer) Do(input any, output func(any)) {
	//    TODO: combine the fields of the incoming contracts.Article with the article template (provided via the constructor),
	//    replace: {{Title}} with contracts.Article.Title
	//             {{Slug}} with contracts.Article.Slug
	//             {{Date}} with contracts.Article.Date.Format("January 2, 2006")
	//             {{Body}} with contracts.Article.Body

	switch input := input.(type) {
	case contracts.Article:

		replacedTemplate := this.template
		replacedTemplate = strings.Replace(replacedTemplate, "{{Title}}", input.Title, -1)
		replacedTemplate = strings.Replace(replacedTemplate, "{{Slug}}", input.Slug, -1)
		replacedTemplate = strings.Replace(replacedTemplate, "{{Date}}", input.Date.Format("January 2, 2006"), -1)
		replacedTemplate = strings.Replace(replacedTemplate, "{{Body}}", input.Body, -1)

		output(contracts.Page{
			Path:    input.Slug,
			Content: replacedTemplate,
		})
	default:
		output(input)
	}
}
