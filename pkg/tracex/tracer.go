package tracex

import (
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func newTraceResource(serviceName string) (*resource.Resource, error) {
	attrs := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(serviceName),
	)
	return resource.Merge(resource.Default(), attrs)
}

func NewTracerProvider(serviceName string, exporter *jaeger.Exporter) (*trace.TracerProvider, error) {
	source, err := newTraceResource(serviceName)
	if err != nil {
		return nil, err
	}

	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(source),
	), nil
}
