package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rongan/otel-gin-instrument/gin-app/log"
	"github.com/rongan/otel-gin-instrument/gin-app/metrics"
	"github.com/rongan/otel-gin-instrument/gin-app/models"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"google.golang.org/grpc/credentials"
	"os"

	"context"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	serviceName  = os.Getenv("SERVICE_NAME")
	collectorURL = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	insecure     = os.Getenv("INSECURE_MODE")
)

func main() {
	log.Init(&log.Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             "info",
		Format:            "json",
		OutputPaths:       []string{"stdout", "./gin.log"},
	})

	cleanup := initTracer()
	defer cleanup(context.Background())

	provider := metrics.InitMeter()
	defer provider.Shutdown(context.Background())

	meter := provider.Meter("GinApp")
	metrics.GenerateMetrics(meter)

	r := gin.Default()
	r.Use(otelgin.Middleware(serviceName))
	models.InitDB()

	installRouters(r)
	r.Run(":8080")
}

func initTracer() func(context.Context) error {

	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if len(insecure) > 0 {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(collectorURL),
		),
	)
	if err != nil {
		log.Fatalw("error", err)
	}

	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Infow("could not set resources", err)
	}

	//otel.SetTextMapPropagator(b3prop.New())
	//fmt.Println(otel.GetTextMapPropagator().Fields())

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resources),
		),
	)
	return exporter.Shutdown
}
