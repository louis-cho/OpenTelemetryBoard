package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	tracepb "otel-server/gen/opentelemetry/proto/collector/trace/v1"
	"otel-server/internal"
)

func main() {
	lis, err := net.Listen("tcp", ":4317")
	if err != nil {
		log.Fatalf("포트 바인딩 실패: %v", err)
	}

	server := grpc.NewServer()
	tracepb.RegisterTraceServiceServer(server, &internal.TraceReceiver{})

	log.Println("🚀 OTLP gRPC 서버 실행 중 (포트 4317)")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("gRPC 서버 실행 실패: %v", err)
	}
}
