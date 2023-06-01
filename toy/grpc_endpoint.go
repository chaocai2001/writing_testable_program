package toy

import (
	"context"

	"github.com/chaocai2001/writing_testable_program/grpc"
)

type GRPC_Endpoint struct {
	grpc.UnimplementedProcessingServiceServer
	processingService ProcessingService
}

func NewGRPC_Endpoint(processingService ProcessingService) *GRPC_Endpoint {
	return &GRPC_Endpoint{
		processingService: processingService,
	}
}

func (ep *GRPC_Endpoint) Process(ctx context.Context, req *grpc.ProcessingRequest) (*grpc.ProcessingReply, error) {
	token, err := ep.processingService.Process(req.GetData())
	if err != nil {
		return nil, err
	}
	return &grpc.ProcessingReply{Token: token}, nil
}

func (ep *GRPC_Endpoint) Retrive(ctx context.Context, req *grpc.RetrivingRequest) (*grpc.RetrivingReply, error) {
	data, err := ep.processingService.Retrive(req.GetToken())
	if err != nil {
		return nil, err
	}

	return &grpc.RetrivingReply{Data: data}, nil
}
