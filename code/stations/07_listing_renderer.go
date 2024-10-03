package stations

import (
	"sort"
	"strings"

	"tobloggan/code/contracts"
)

type ListingRenderer struct {
	articleListing []contracts.Article
	template       string
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
		return this.articleListing[i].Date.Before(this.articleListing[j].Date)
	})

	var sb strings.Builder
	for i, listing := range this.articleListing {
		if i != 0 {
			sb.WriteString("\n")
		}
		sb.WriteString("<li>")
		sb.WriteString(listing.Title)
		sb.WriteString("<\\li>")
	}
	pageContent := strings.Replace(this.template, "{{Listing}}", sb.String(), 1)

	output(contracts.Page{Path: "/", Content: pageContent})
}
