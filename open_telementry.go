package GolangTechTask

import (
	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func CreateOpenTelemetry(c *Config) (*sdktrace.TracerProvider, error) {
	o := []sdktrace.TracerProviderOption{
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	}
	if c.Trace {
		exporter, err := stdout.New(stdout.WithPrettyPrint())
		if err != nil {
			return nil, err
		}
		o = append(o,
			sdktrace.WithBatcher(exporter),
		)
	}
	tp := sdktrace.NewTracerProvider(o...)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}
