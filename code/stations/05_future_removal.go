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
	switch input := input.(type) {
	case contracts.Article:
		if !input.Date.After(this.current) {
			output(input)
		}
	default:
		output(input)
	}

}
