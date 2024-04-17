package main

import (
	"airport-app-backend/config"
	"airport-app-backend/server"
	"os"
	"time"

	"contrib.go.opencensus.io/exporter/jaeger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

var Database *gorm.DB

func initTracing() {
	exporter, err := jaeger.NewExporter(jaeger.Options{
		CollectorEndpoint: "http://localhost:14268/api/traces",
		ServiceName:       "airport-service",
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create Jaeger exporter")
	}

	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}

func main() {
	initTracing()
	// Gin set mode release / debug
	if config.ProductionMode {
		gin.SetMode(gin.ReleaseMode)
	}
	// Global rs/zerolog config
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    false,
		TimeFormat: time.RFC1123Z,
	})
	runServer()
}

func runServer() {
	log.Info().Msg("Starting application")

	srv := server.NewServer(gin.New())

	// Application startup
	err := srv.RunServer()
	if err != nil {
		log.Fatal().Err(err).Msg("Application startup failed")
	}
}
