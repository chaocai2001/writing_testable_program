package untestable_test

import (
	"context"
	"errors"
	"testing"

	"github.com/chaocai2001/writing_testable_program/grpc"
	"github.com/chaocai2001/writing_testable_program/storage"
	"github.com/chaocai2001/writing_testable_program/untestable"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type mockRedisClient struct {
	storage.MockRedisClientInterface
	data map[string]string
}

func (m *mockRedisClient) StoreData(key, value string) error {
	m.data[key] = value
	return nil
}

func (m *mockRedisClient) RetiveData(key string) (string, error) {
	value, ok := m.data[key]
	if !ok {
		return "", errors.New("data not found")
	}
	return value, nil
}

func TestProcessService_Process(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := &mockRedisClient{data: make(map[string]string)}
	ps := &untestable.ProcessService{storage: mockStorage}

	ctx := context.Background()
	req := &grpc.ProcessingRequest{Data: "test"}

	reply, err := ps.Process(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, reply.Token)

	retriveReq := &grpc.RetrivingRequest{Token: reply.Token}
	retriveReply, err := ps.Retrive(ctx, retriveReq)
	assert.NoError(t, err)
	assert.Equal(t, "tset", retriveReply.Data)
}

func TestProcessService_Process_EmptyInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := &mockRedisClient{data: make(map[string]string)}
	ps := &untestable.ProcessService{storage: mockStorage}

	ctx := context.Background()
	req := &grpc.ProcessingRequest{Data: ""}

	_, err := ps.Process(ctx, req)
	assert.Error(t, err)
	assert.Equal(t, "Invalid input", err.Error())
}
