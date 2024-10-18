package main

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	api "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	res, err := resource.New(ctx, resource.WithAttributes(attribute.String("service.name", "service")))
	if err != nil {
		log.Println(err.Error())
	}

	go serveMetrics()

	exp, err := prometheus.New()
	if err != nil {
		log.Println(err.Error())
	}

	provider := metric.NewMeterProvider(
		metric.WithReader(exp),
		metric.WithResource(res),
	)
	otel.SetMeterProvider(provider)
	opts := api.WithAttributes(
		attribute.Key("grpc_method").String("Get"),
		attribute.Key("grpc_method").String("Describe"),
		attribute.Key("grpc_method").String("List"),
	)

	meter := provider.Meter("METER NAME")
	histogram, err := meter.Float64Histogram(
		"HISTOGRAM NAME",
		api.WithDescription("HISTOGRAM DESCRIPTION"),
		api.WithExplicitBucketBoundaries(
			0, 0.0005, 0.001, 0.005, 0.01, 0.015, 0.02, 0.025,
			0.03, 0.035, 0.04, 0.05, 0.06, 0.07, 0.08, 0.09,
			0.1, 0.13, 0.16, 0.2, 0.25, 0.3, 0.35, 0.4, 0.45,
			0.5, 0.55, 0.6, 0.65, 0.7, 0.8, 0.9, 1,
			1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5,
			5.5, 6, 7, 8, 9, 10, 12, 15,
			17, 20, 50, 100),
	)
	if err != nil {
		log.Println(err.Error())
	}

	histogram.Record(ctx, 50, opts)
	histogram.Record(ctx, 75, opts)
	histogram.Record(ctx, 300, opts)
	NewServerM

	<-ctx.Done()
}

func serveMetrics() {
	log.Println("serving metrics at localhost:8080/metrics")
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("error serving http: %v\n", err)
	}
}

