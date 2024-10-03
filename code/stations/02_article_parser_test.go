package stations

import (
	"testing"
	"time"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func (this *ArticleParserFixture) Setup() {
	this.station = NewArticleParser()
}

func TestArticleParserFixture(t *testing.T) {
	gunit.Run(new(ArticleParserFixture), t)
}

type ArticleParserFixture struct {
	StationFixture
	markdownErr error
}

func (this *ArticleParserFixture) TestArticleMetaAndContentReadFromDiskAndEmitted() {
	this.do(contracts.SourceFile(article1Content))
	this.So(this.outputs, should.Equal, []any{
		contracts.Article{
			Slug:  "/article/1",
			Title: "Article 1",
			Date:  time.Date(2024, time.September, 4, 0, 0, 0, 0, time.UTC),
			Body:  "The contents of article 1.",
		},
	})
}
func (this *ArticleParserFixture) TestMissingDivider() {
	this.do(contracts.SourceFile(article2Content))
	this.So(this.outputs[0], should.Wrap, errMalformedContent)
}
func (this *ArticleParserFixture) TestMalformedMetadata() {
	this.do(contracts.SourceFile(article3Content))
	this.So(this.outputs[0], should.Wrap, errMalformedContent)
}

const article1Content = `{
	"date": "2024-09-04T00:00:00Z",
	"slug": "/article/1",
	"title": "Article 1"
}

+++

The contents of article 1.`
const article2Content = `{}

The contents of article 2.`
const article3Content = `{"gumbo""gumbo"}

+++

The contents of article 3.`
