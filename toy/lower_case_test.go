package toy

import (
	"testing"
)

func TestProcessing(t *testing.T) {
	req := "Hello World!"
	processor := NewLowerCaseProcessor()
	rep, err := processor.Process(req)
	if err != nil {
		t.Error(err)
		return
	}
	if rep != "hello world!" {
		t.Errorf("the expected value is 'hello world!', but the actual value is %s", rep)
		return
	}
}
