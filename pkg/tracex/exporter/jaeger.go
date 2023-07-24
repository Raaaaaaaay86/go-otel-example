package exporter

import "go.opentelemetry.io/otel/exporters/jaeger"

func NewJaegerExporter(endpoint string) (*jaeger.Exporter, error) {
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
}
