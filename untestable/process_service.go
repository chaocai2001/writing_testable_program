package untestable

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"

	"github.com/chaocai2001/writing_testable_program/grpc"
	"github.com/chaocai2001/writing_testable_program/storage"
)

type ProcessService struct {
	storage *storage.RedisClient
	grpc.UnimplementedProcessingServiceServer
}

func NewProcessService() *ProcessService {
	return &ProcessService{
		storage: storage.NewRedisClient(),
	}
}

func (ps *ProcessService) process(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", errors.New("Invalid input")
	}
	processed := ""
	for _, r := range input {
		processed = string(r) + processed
	}
	return processed, nil
}

func (ps *ProcessService) generateToken() string {
	id, err := uuid.NewUUID()
	if err != nil {
		panic("failed to generate ID")
	}
	return id.String()
}

func (ps *ProcessService) Process(ctx context.Context, req *grpc.ProcessingRequest) (*grpc.ProcessingReply, error) {
	processed, err := ps.process(req.GetData())
	if err != nil {
		return nil, err
	}
	token := ps.generateToken()
	ps.storage.StoreData(token, processed)
	return &grpc.ProcessingReply{Token: token}, nil
}

func (ps *ProcessService) Retrive(ctx context.Context, req *grpc.RetrivingRequest) (*grpc.RetrivingReply, error) {
	data, err := ps.storage.RetiveData(req.GetToken())
	if err != nil {
		return nil, err
	}

	return &grpc.RetrivingReply{Data: data}, nil
}
