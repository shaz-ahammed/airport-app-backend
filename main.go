package main

import (
    "os"
    "time"
    "net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
    // Initialize Jaeger exporter and tracing
    initTracing()

    router := gin.New()

    router.GET("/books", handleGetBooks)

    server := &http.Server{
        Addr:         ":8080",
        Handler:      router,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    // Start HTTP server
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal().Err(err).Msg("Failed to start HTTP server")
        }
    }()

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
