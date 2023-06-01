package toy_test

import (
	"errors"
	"testing"

	"github.com/chaocai2001/writing_testable_program"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) RetiveData(token string) (string, error) {
	args := m.Called(token)
	return args.String(0), args.Error(1)
}

func (m *MockStorage) StoreData(token string, data string) error {
	args := m.Called(token, data)
	return args.Error(0)
}

type MockProcessor struct {
	mock.Mock
}

func (m *MockProcessor) Process(raw string) (string, error) {
	args := m.Called(raw)
	return args.String(0), args.Error(1)
}

type MockTokenCreator struct {
	mock.Mock
}

func (m *MockTokenCreator) CreateToken(data string) string {
	args := m.Called(data)
	return args.String(0)
}

func TestProcessingService_Process(t *testing.T) {
	mockStorage := new(MockStorage)
	mockProcessor := new(MockProcessor)
	mockTokenCreator := new(MockTokenCreator)

	processingService := toy.NewProcessingService(mockProcessor, mockTokenCreator, mockStorage)

	raw := "raw_data"
	processed := "processed_data"
	token := "token"

	mockProcessor.On("Process", raw).Return(processed, nil)
	mockTokenCreator.On("CreateToken", processed).Return(token)
	mockStorage.On("StoreData", token, processed).Return(nil)

	resultToken, err := processingService.Process(raw)

	assert.NoError(t, err)
	assert.Equal(t, token, resultToken)

	mockProcessor.AssertExpectations(t)
	mockTokenCreator.AssertExpectations(t)
	mockStorage.AssertExpectations(t)
}

func TestProcessingService_Retrive(t *testing.T) {
	mockStorage := new(MockStorage)
	mockProcessor := new(MockProcessor)
	mockTokenCreator := new(MockTokenCreator)

	processingService := toy.NewProcessingService(mockProcessor, mockTokenCreator, mockStorage)

	token := "token"
	data := "data"

	mockStorage.On("RetiveData", token).Return(data, nil)

	resultData, err := processingService.Retrive(token)

	assert.NoError(t, err)
	assert.Equal(t, data, resultData)

	mockStorage.AssertExpectations(t)
}

func TestProcessingService_Process_Error(t *testing.T) {
	mockStorage := new(MockStorage)
	mockProcessor := new(MockProcessor)
	mockTokenCreator := new(MockTokenCreator)

	processingService := toy.NewProcessingService(mockProcessor, mockTokenCreator, mockStorage)

	raw := "raw_data"

	mockProcessor.On("Process", raw).Return("", errors.New("processing error"))

	_, err := processingService.Process(raw)

	assert.Error(t, err)

	mockProcessor.AssertExpectations(t)
}
