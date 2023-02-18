package toy

import (
	"log"
	"time"
)

type ProcessorLogDecorator struct {
	processor Processor
}

func (pld *ProcessorLogDecorator) Process(raw string) (string, error) {
	log.Printf("[info] the request is %s", raw)
	ret, err := pld.processor.Process(raw)
	log.Printf("[info] the ouput is %s %v", ret, err)
	return ret, err
}

type ProcessorTimerDecorator struct {
	processor Processor
}

func (ptd *ProcessorTimerDecorator) Process(raw string) (string, error) {
	startTime := time.Now()
	ret, err := ptd.processor.Process(raw)
	duration := time.Since(startTime).Milliseconds()
	log.Printf("[info] the time spent %d ms", duration)
	return ret, err
}
