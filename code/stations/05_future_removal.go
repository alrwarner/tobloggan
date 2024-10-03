package stations

import (
	"time"

	"tobloggan/code/contracts"
)

type FutureRemoval struct {
	current time.Time
}

func NewFutureRemoval(now time.Time) *FutureRemoval {
	return &FutureRemoval{current: now}
}

func (this *FutureRemoval) Do(input any, output func(any)) {
	//    TODO: given a contracts.Article, only output it if the Date field is not after the current time (passed into constructor).

	switch input := input.(type) {
	case contracts.Article:
		if !input.Date.After(this.current) {
			output(input)
		}
	default:
		output(input)
	}

}
