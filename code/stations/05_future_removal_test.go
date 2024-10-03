package stations

//func (this *FutureRemovalFixture) TestPastArticleKept() {}
//func (this *FutureRemovalFixture) TestCurrentArticleKept() {}
//func (this *FutureRemovalFixture) TestFutureArticleDropped() {}

import (
	"testing"
	"time"
	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestFutureRemovalFixture(t *testing.T) {
	gunit.Run(new(FutureRemovalFixture), t)
}

type FutureRemovalFixture struct {
	StationFixture
}

func (this *FutureRemovalFixture) Setup() {
	this.station = NewFutureRemoval(time.Now())
}

func (this *FutureRemovalFixture) TestPastArticleKept() {
	article := contracts.Article{
		Draft: true,
		Slug:  "",
		Title: "",
		Date:  time.Now().Add(-24 * time.Hour),
		Body:  "# This is some past content",
	}
	this.station.Do(article, this.output)
	this.So(this.outputs, should.HaveLength, 1)
}

func (this *FutureRemovalFixture) TestCurrentArticleKept() {
	article := contracts.Article{
		Draft: false,
		Slug:  "",
		Title: "",
		Date:  time.Now().Add(-1 * time.Minute),
		Body:  "# This is some current content",
	}
	this.station.Do(article, this.output)
	this.So(this.outputs, should.HaveLength, 1)
}

func (this *FutureRemovalFixture) TestFutureArticleDropped() {
	article := contracts.Article{
		Draft: false,
		Slug:  "",
		Title: "",
		Date:  (time.Now().Add(5 * time.Hour)),
		Body:  "# This is some future content",
	}
	this.station.Do(article, this.output)
	this.So(this.outputs, should.HaveLength, 0)
}
