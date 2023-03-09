package toy

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	ggrpc "google.golang.org/grpc"

	"github.com/chaocai2001/writing_testable_program/grpc"
)

const Port = 8088

func createProcessingService() *ProcessingService {
	processor := NewLowerCaseProcessor()
	tokenCreator := &MockTokenCreator{}
	storage := NewLocalMapStore()
	return NewProcessingService(processor, tokenCreator, storage)
}

func createProcessingServiceWithDecorator() *ProcessingService {
	processor := &ProcessorLogDecorator{
		&ProcessorTimerDecorator{NewLowerCaseProcessor()},
	}
	tokenCreator := &MockTokenCreator{}
	storage := NewLocalMapStore()
	return NewProcessingService(processor, tokenCreator, storage)
}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}

func startServer(t *testing.T) {
	// endPointImpl := NewGRPC_Endpoint(createProcessingService())
	endPointImpl := NewGRPC_Endpoint(createProcessingServiceWithDecorator())
	gRPCServer := ggrpc.NewServer()
	grpc.RegisterProcessingServiceServer(gRPCServer, endPointImpl)
	log.Printf("Processing server is started ...")
	if err := gRPCServer.Serve(getNetListener(Port)); err != nil {
		t.Fatal(err)
	}
}

func runClient(t *testing.T) {
	addr := fmt.Sprintf("localhost:%d", Port)
	conn, err := ggrpc.Dial(addr, ggrpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := grpc.NewProcessingServiceClient(conn)
	reply, err1 := client.Process(context.TODO(), &grpc.ProcessingRequest{Data: "Hello World"})
	if err1 != nil {
		t.Error(err1)
		return
	}
	token := reply.GetToken()
	fmt.Println(token)
	rreply, err2 := client.Retrive(context.TODO(), &grpc.RetrivingRequest{Token: token})
	if err2 != nil {
		t.Error(err2)
		return
	}
	if rreply.GetData() != "hello world" {
		t.Errorf("the expected value is 'hello world', but the actual value is %s", rreply)
		return
	}
}

func TestBasicEndpoint(t *testing.T) {
	go startServer(t)
	runClient(t)
	time.Sleep(time.Second * 1)
}
