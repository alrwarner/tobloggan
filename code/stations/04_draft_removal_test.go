package stations

import (
	"testing"
	"time"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestDraftRemovalFixture(t *testing.T) {
	gunit.Run(new(DraftRemovalFixture), t)
}

type DraftRemovalFixture struct {
	StationFixture
}

func (this *DraftRemovalFixture) Setup() {
	this.station = NewDraftRemoval()
}

func (this *DraftRemovalFixture) TestDraftDropped() {
	article := contracts.Article{
		Draft: true,
		Slug:  "",
		Title: "",
		Date:  time.Time{},
		Body:  "# This is some content",
	}
	this.station.Do(article, this.output)
	this.So(this.outputs, should.HaveLength, 0)
}

func (this *DraftRemovalFixture) TestNonDraftRetained() {
	article := contracts.Article{
		Draft: false,
		Slug:  "",
		Title: "",
		Date:  time.Time{},
		Body:  "# This is some content",
	}
	this.station.Do(article, this.output)
	this.So(this.outputs, should.HaveLength, 1)
}

//func (this *DraftRemovalFixture) TestDraftDropped() {}
//func (this *DraftRemovalFixture) TestNonDraftRetained() {}
