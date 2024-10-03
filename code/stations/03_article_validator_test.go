package stations

import (
	"testing"
	"time"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func (this *ArticleValidatorFixture) Setup() {
	this.station = NewArticleValidator()
}

func TestArticleValidatorFixture(t *testing.T) {
	gunit.Run(new(ArticleParserFixture), t)
}

type ArticleValidatorFixture struct {
	StationFixture
	markdownErr error
}

func (this *ArticleValidatorFixture) TestValidArticle() {
	article := contracts.Article{
		Draft: false,
		Slug:  "/article/1",
		Title: "Article 1",
		Date:  time.Date(2024, time.September, 4, 0, 0, 0, 0, time.UTC),
		Body:  "The contents of article 1.",
	}
	this.do(article)
	this.So(this.outputs, should.Equal, []any{
		article,
	})
}
func (this *ArticleValidatorFixture) TestInvalidSlugs() {
	article := contracts.Article{
		Draft: false,
		Slug:  "/article/ / /",
		Title: "Article 1",
		Date:  time.Date(2024, time.September, 4, 0, 0, 0, 0, time.UTC),
		Body:  "The contents of article 1.",
	}
	this.do(article)
	this.So(this.outputs, should.Wrap, errMalformedContent)
}
func (this *ArticleValidatorFixture) TestInvalidTitles() {
	article := contracts.Article{
		Draft: false,
		Slug:  "/article/1/",
		Title: "",
		Date:  time.Date(2024, time.September, 4, 0, 0, 0, 0, time.UTC),
		Body:  "The contents of article 1.",
	}
	this.do(article)
	this.So(this.outputs, should.Wrap, errMalformedContent)

	article.Title = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	this.do(article)
	this.So(this.outputs, should.Wrap, errMalformedContent)
}
func (this *ArticleValidatorFixture) TestSlugsMustBeUnique() {
	article := contracts.Article{
		Draft: false,
		Slug:  "/article/1/",
		Title: "",
		Date:  time.Date(2024, time.September, 4, 0, 0, 0, 0, time.UTC),
		Body:  "The contents of article 1.",
	}
	this.do(article)
	this.do(article)
	this.So(this.outputs, should.Wrap, errDuplicateSlug)
}
