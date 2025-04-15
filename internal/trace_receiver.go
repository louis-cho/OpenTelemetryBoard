package internal

import (
	"context"
	"encoding/json"
	"log"

	tracepb "otel-server/gen/opentelemetry/proto/collector/trace/v1"
)

type TraceReceiver struct {
	tracepb.UnimplementedTraceServiceServer
}

func (t *TraceReceiver) Export(ctx context.Context, req *tracepb.ExportTraceServiceRequest) (*tracepb.ExportTraceServiceResponse, error) {
	for _, resourceSpan := range req.ResourceSpans {
		jsonBytes, _ := json.MarshalIndent(resourceSpan, "", "  ")
		log.Printf("🟢 Trace 수신:\n%s\n", string(jsonBytes))
	}
	return &tracepb.ExportTraceServiceResponse{}, nil
}
