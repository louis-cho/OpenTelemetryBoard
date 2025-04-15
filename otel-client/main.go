package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	// OTLP Exporter 생성
	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint("localhost:4317"),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
	)
	if err != nil {
		log.Fatalf("❌ OTLP exporter 생성 실패: %v", err)
	}

	// TracerProvider 생성
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("client-test"),
		)),
	)
	otel.SetTracerProvider(tp)

	// Tracer 만들고 span 시작
	tracer := otel.Tracer("client-tracer")
	ctx, span := tracer.Start(ctx, "client-span-1")
	time.Sleep(100 * time.Millisecond)
	span.End()

	log.Println("✅ Trace 전송 완료")

	// 버퍼에 쌓인 trace flush
	if err := tp.Shutdown(ctx); err != nil {
		log.Fatalf("❌ TraceProvider 종료 실패: %v", err)
	}
}
