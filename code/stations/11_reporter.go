package stations

import (
	"sync/atomic"

	"tobloggan/code/contracts"
)

type Reporter struct {
	logger contracts.Logger
	failed *atomic.Bool
}

func NewReporter(logger contracts.Logger, failed *atomic.Bool) contracts.Station {
	return &Reporter{logger: logger, failed: failed}
}

func (this *Reporter) Do(input any, output func(any)) {
	switch input := input.(type) {
	case error:
		this.failed.Store(true)
		this.logger.Printf("err: %v", input)
	case contracts.Page:
		this.logger.Printf("page: %s", input.Path)
	default:
		output(input)
	}
}
