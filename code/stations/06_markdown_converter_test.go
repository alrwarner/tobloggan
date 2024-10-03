package stations

import (
	"testing"
	"time"

	"tobloggan/code/contracts"
	"tobloggan/code/markdown"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestMarkdownConverterFixture(t *testing.T) {
	gunit.Run(new(MarkdownConverterFixture), t)
}

type MarkdownConverterFixture struct {
	StationFixture
}

func (this *MarkdownConverterFixture) Setup() {
	this.station = NewMarkdownConverter(markdown.NewConverter())
}

func (this *MarkdownConverterFixture) TestConverter() {
	article := contracts.Article{
		Draft: false,
		Slug:  "",
		Title: "",
		Date:  time.Time{},
		Body:  "# This is some content",
	}
	this.station.Do(article, this.output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0].(contracts.Article).Body, should.Equal, "<h1>This is some content</h1>\n")
	}
}
