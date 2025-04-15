package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	collectortrace "go.opentelemetry.io/proto/otlp/collector/trace/v1"
)

type MyTraceServer struct {
	collectortrace.UnimplementedTraceServiceServer
}

func (s *MyTraceServer) Export(ctx context.Context, req *collectortrace.ExportTraceServiceRequest) (*collectortrace.ExportTraceServiceResponse, error) {
	fmt.Println("🟢 Trace received!")
	for _, span := range req.ResourceSpans {
		fmt.Printf("Resource: %+v\n", span.Resource.Attributes)
	}
	return &collectortrace.ExportTraceServiceResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4317")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	collectortrace.RegisterTraceServiceServer(grpcServer, &MyTraceServer{})

	fmt.Println("🚀 gRPC OTLP 서버 실행 중 (localhost:4317)")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC 서버 종료: %v", err)
	}
}
