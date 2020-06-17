package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-app-starter/middleware"
	"os"
	"runtime"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Global zerolog config
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    false,
		TimeFormat: time.RFC1123Z,
	})

	// Gin router and middleware config
	router := gin.New()
	setupGinMiddleware(router)

	// Example ping request.
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	log.Info().
		Int("pid", os.Getpid()).
		Int("parent-pid", os.Getppid()).
		Int("num-cpu", runtime.NumCPU()).
		Int("GOMAXPROCS", runtime.GOMAXPROCS(-1)).
		Uint64("total-memory-obtained-from-sys-MB", memStats.Sys/1024/1024).
		Uint64("total-memory-allocated-heap-MB", memStats.TotalAlloc/1024/1024).
		Msg("Starting application")

	// Application startup

	// For this app to work on Heroku we need to read PORT value from env variable. For local it will be set to predefined value.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(fmt.Sprintf("0.0.0.0:%s", port)); err != nil {
		log.Fatal().Err(err).Msg("Startup failed")
	}
}

func setupGinMiddleware(ginRouter *gin.Engine) {
	log.Info().Msg("Configuring GIN middleware")
	ginRouter.Use(gin.Recovery()) // Default recovery middleware

	ginRouter.Use(middleware.DisableCache())
	ginRouter.Use(middleware.AddSecurityHeaders(true))

	if detailedLogging := os.Getenv("DETAILED_LOGGING"); detailedLogging != "false" {
		ginRouter.Use(middleware.ZerologConsoleRequestLogging())
	}
}
