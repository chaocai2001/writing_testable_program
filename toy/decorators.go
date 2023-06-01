package toy

import (
	"log"
	"time"
)

type ProcessorLogDecorator struct {
	Processor Processor
}

func (pld *ProcessorLogDecorator) Process(raw string) (string, error) {
	log.Printf("[info] the request is %s", raw)
	ret, err := pld.Processor.Process(raw)
	log.Printf("[info] the ouput is %s %v", ret, err)
	return ret, err
}

type ProcessorTimerDecorator struct {
	Processor Processor
}

func (ptd *ProcessorTimerDecorator) Process(raw string) (string, error) {
	startTime := time.Now()
	ret, err := ptd.Processor.Process(raw)
	duration := time.Since(startTime).Milliseconds()
	log.Printf("[info] the time spent %d ms", duration)
	return ret, err
}

// Decorator pattern can be implemented by functional programming
type FN func(data string) (string, error)

func DecoratorFn(f FN) FN {
	return func(data string) (string, error) {
		log.Printf("invoke by %s", data)
		ret, err := f(data)
		log.Printf("return with %s, %v", ret, err)
		return ret, err
	}
}
