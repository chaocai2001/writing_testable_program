package toy_test

import (
	"testing"

	"github.com/chaocai2001/writing_testable_program"
)

type MockStorage struct{}
type MockProcessor struct{}
type MockTokenCreator struct{}

func (ms *MockStorage) RetiveData(token string) (string, error) {
	return "mockData", nil
}

func (ms *MockStorage) StoreData(token string, data string) error {
	return nil
}

func (mp *MockProcessor) Process(raw string) (string, error) {
	return "mockProcessed", nil
}

func (mtc *MockTokenCreator) CreateToken(data string) string {
	return "mockToken"
}

func TestProcessingService_Process(t *testing.T) {
	mockStorage := &MockStorage{}
	mockProcessor := &MockProcessor{}
	mockTokenCreator := &MockTokenCreator{}

	processingService := toy.NewProcessingService(mockProcessor, mockTokenCreator, mockStorage)

	token, err := processingService.Process("rawData")
	if err != nil {
		t.Errorf("Process() error = %v", err)
		return
	}

	if token != "mockToken" {
		t.Errorf("Process() token = %v, want %v", token, "mockToken")
	}
}

func TestProcessingService_Retrive(t *testing.T) {
	mockStorage := &MockStorage{}
	mockProcessor := &MockProcessor{}
	mockTokenCreator := &MockTokenCreator{}

	processingService := toy.NewProcessingService(mockProcessor, mockTokenCreator, mockStorage)

	data, err := processingService.Retrive("mockToken")
	if err != nil {
		t.Errorf("Retrive() error = %v", err)
		return
	}

	if data != "mockData" {
		t.Errorf("Retrive() data = %v, want %v", data, "mockData")
	}
}
