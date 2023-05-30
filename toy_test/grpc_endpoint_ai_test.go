package toy_test

import (
	"context"
	"errors"
	"testing"

	"github.com/chaocai2001/writing_testable_program"
	"github.com/chaocai2001/writing_testable_program/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProcessingService struct {
	mock.Mock
}

func (m *MockProcessingService) Process(data string) (string, error) {
	args := m.Called(data)
	return args.String(0), args.Error(1)
}

func (m *MockProcessingService) Retrive(token string) (string, error) {
	args := m.Called(token)
	return args.String(0), args.Error(1)
}

func TestProcess(t *testing.T) {
	mockService := new(MockProcessingService)
	mockService.On("Process", "test_data").Return("test_token", nil)

	endpoint := toy.NewGRPC_Endpoint(mockService)
	req := &grpc.ProcessingRequest{Data: "test_data"}

	reply, err := endpoint.Process(context.Background(), req)

	assert.Nil(t, err)
	assert.Equal(t, "test_token", reply.Token)
	mockService.AssertExpectations(t)
}

func TestProcess_Error(t *testing.T) {
	mockService := new(MockProcessingService)
	mockService.On("Process", "test_data").Return("", errors.New("test_error"))

	endpoint := toy.NewGRPC_Endpoint(mockService)
	req := &grpc.ProcessingRequest{Data: "test_data"}

	reply, err := endpoint.Process(context.Background(), req)

	assert.NotNil(t, err)
	assert.Nil(t, reply)
	mockService.AssertExpectations(t)
}

func TestRetrive(t *testing.T) {
	mockService := new(MockProcessingService)
	mockService.On("Retrive", "test_token").Return("test_data", nil)

	endpoint := toy.NewGRPC_Endpoint(mockService)
	req := &grpc.RetrivingRequest{Token: "test_token"}

	reply, err := endpoint.Retrive(context.Background(), req)

	assert.Nil(t, err)
	assert.Equal(t, "test_data", reply.Data)
	mockService.AssertExpectations(t)
}

func TestRetrive_Error(t *testing.T) {
	mockService := new(MockProcessingService)
	mockService.On("Retrive", "test_token").Return("", errors.New("test_error"))

	endpoint := toy.NewGRPC_Endpoint(mockService)
	req := &grpc.RetrivingRequest{Token: "test_token"}

	reply, err := endpoint.Retrive(context.Background(), req)

	assert.NotNil(t, err)
	assert.Nil(t, reply)
	mockService.AssertExpectations(t)
}
