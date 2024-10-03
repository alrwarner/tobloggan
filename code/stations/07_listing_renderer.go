package stations

import (
	"fmt"
	"sort"
	"strings"

	"tobloggan/code/contracts"
)

type ListingRenderer struct {
	articleListing []contracts.Article
	template       string
}

func NewListingRenderer(listingTemplate string) *ListingRenderer {
	return &ListingRenderer{
		articleListing: make([]contracts.Article, 0),
		template:       listingTemplate,
	}
}

func (this *ListingRenderer) Do(input any, output func(any)) {
	//    TODO: given a contracts.Article, append it to a slice and send it on
	switch input := input.(type) {
	case contracts.Article:
		this.articleListing = append(this.articleListing, input)
		output(input)
	default:
		output(input)
	}

}
func (this *ListingRenderer) Finalize(output func(any)) {
	//    TODO: sort the slice (by Date), generate a <li> for each article in a big string,
	//    TODO: pageContent := strings.Replace(this.template, "{{Listing}}", renderedListing, 1)
	//    TODO: output(contracts.Page{Path: "/", Content: pageContent})

	sort.Slice(this.articleListing, func(i, j int) bool {
		return this.articleListing[i].Date.After(this.articleListing[j].Date)
	})

	var sb strings.Builder
	for i, listing := range this.articleListing {
		if i != 0 {
			sb.WriteString("\n")
		}

		sb.WriteString(`<li>`)
		// <a href=SLUG>TITLE<a>
		sb.WriteString(fmt.Sprintf(`<a href="%s">%s<\a>`, listing.Slug, listing.Title))
		sb.WriteString(`<\li>`)
	}
	pageContent := strings.Replace(this.template, "{{Listing}}", sb.String(), 1)

	output(contracts.Page{Path: "/", Content: pageContent})
}
