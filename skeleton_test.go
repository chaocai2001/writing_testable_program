package toy

import (
	"strings"
	"testing"
)

type MockTokenCreator struct {
}

func (mtc *MockTokenCreator) CreateToken(data string) string {
	return strings.ToUpper(data)
}

func TestBasicLogic(t *testing.T) {
	sample := "Hello World"
	processor := NewLowerCaseProcessor()
	tokenCreator := &MockTokenCreator{}
	storage := NewLocalMapStore()
	processingService := NewProcessingService(processor, tokenCreator, storage)
	token, err := processingService.Process(sample)
	if err != nil {
		t.Error(err)
		return
	}
	retivedData, err1 := processingService.Retrive(token)
	if err1 != nil {
		t.Error(err)
		return
	}
	if retivedData != "hello world" {
		t.Errorf("the expected value is 'hello world', but the actual value is %s", retivedData)
		return
	}
}
