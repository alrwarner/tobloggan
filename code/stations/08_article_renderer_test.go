package stations

import (
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
	"testing"
	"time"
	"tobloggan/code/contracts"
	"tobloggan/code/html"
)

func TestArticleRendererFixture(t *testing.T) {
	gunit.Run(new(ArticleRendererFixture), t)
}

type ArticleRendererFixture struct {
	StationFixture
}

func (this *ArticleRendererFixture) Setup() {
	this.station = NewArticleRenderer(html.ArticleTemplate)
}

func (this *ArticleRendererFixture) TestRendering() {
	article := contracts.Article{
		Draft: false,
		Slug:  "/the-test-path/of-the/article",
		Title: "Bryan is the Best!",
		Date:  time.Now(),
		Body:  "Bryan is a famous author and **chicken master**!!",
	}
	this.station.Do(article, this.output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0].(contracts.Page).Path, should.Equal, "/the-test-path/of-the/article")
	}
}
