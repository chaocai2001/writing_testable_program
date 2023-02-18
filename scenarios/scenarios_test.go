package scenarios

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"testing"

	//	"time"

	. "github.com/chaocai2001/writing_testable_program"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc/credentials/insecure"

	ggrpc "google.golang.org/grpc"

	"github.com/chaocai2001/writing_testable_program/grpc"
)

type MockTokenCreator struct {
}

func (mtc *MockTokenCreator) CreateToken(data string) string {
	return strings.ToUpper(data)
}

const Port = 8082

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
	endPointImpl := NewGRPC_Endpoint(createProcessingServiceWithDecorator())
	gRPCServer := ggrpc.NewServer()
	grpc.RegisterProcessingServiceServer(gRPCServer, endPointImpl)
	log.Printf("Processing server is started ...")
	if err := gRPCServer.Serve(getNetListener(Port)); err != nil {
		t.Fatal(err)
	}
}

func TestHappyScenario(t *testing.T) {
	go startServer(t)
	var token string
	Convey("Scenario: process the data", t, func() {
		Convey("Given a connected service client\n", func() {
			addr := fmt.Sprintf("localhost:%d", Port)
			conn, err := ggrpc.Dial(addr, ggrpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				t.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			client := grpc.NewProcessingServiceClient(conn)
			Convey("When send the data to the service\n", func() {

				reply, err1 := client.Process(context.TODO(), &grpc.ProcessingRequest{Data: "Hello World"})
				if err1 != nil {
					t.Error(err1)
					return
				}
				token = reply.GetToken()

				Convey("Then get a token\n", func() {

					So(len(token), ShouldBeGreaterThan, 0)
				})
			})
			Convey("When retrive the data by the token\n", func() {
				reply, err1 := client.Retrive(context.TODO(), &grpc.RetrivingRequest{Token: token})
				if err1 != nil {
					t.Error(err1)
					return
				}
				data := reply.GetData()
				Convey("Then get the processed data", func() {
					So(data, ShouldEqual, "hello world")
				})
			})
		})
	})
}
