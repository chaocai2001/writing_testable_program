package toy

import (
	"strings"
	"time"
)

type LowerCaseProcessor struct {
}

func NewLowerCaseProcessor() Processor {
	return &LowerCaseProcessor{}
}

func (p *LowerCaseProcessor) Process(raw string) (string, error) {
	time.Sleep(time.Millisecond * 200)
	return strings.ToLower(raw), nil
}
