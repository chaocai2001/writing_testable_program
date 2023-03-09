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

type testCase struct {
	input  string
	expect string
}

func TestBasicLogic(t *testing.T) {
	testcases := []testCase{
		testCase{
			"Hello World", "hello world",
		},
		testCase{
			"Hello World1", "hello world1",
		},
	}

	processor := NewLowerCaseProcessor()
	tokenCreator := &MockTokenCreator{}
	storage := NewLocalMapStore()
	processingService := NewProcessingService(processor, tokenCreator, storage)
	for _, testcase := range testcases {
		token, err := processingService.Process(testcase.input)
		if err != nil {
			t.Error(err)
			return
		}
		retivedData, err1 := processingService.Retrive(token)
		if err1 != nil {
			t.Error(err)
			return
		}
		if retivedData != testcase.expect {
			t.Errorf("the expected value is 'hello world', but the actual value is %s", retivedData)
			return
		}
	}
}
