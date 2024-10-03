package stations

import (
	"time"

	"tobloggan/code/contracts"
)

type FutureRemoval struct{}

func (this *FutureRemoval) Do(input any, output func(any)) {
	//    TODO: given a contracts.Article, only output it if the Date field is not after the current time (passed into constructor).

	switch input := input.(type) {
	case contracts.Article:
		if !input.Date.After(time.Now()) {
			output(input)
		}
	default:
		output(input)
	}

}
