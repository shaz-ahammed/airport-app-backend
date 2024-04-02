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

	log.Info().
		Int("pid", os.Getpid()).
		Int("parent-pid", os.Getppid()).
		Int("num-cpu", runtime.NumCPU()).
		Int("GOMAXPROCS", runtime.GOMAXPROCS(-1)).
		Msg("Starting application")

	// Application startup
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatal().Err(err).Msg("Startup failed")
	}
}

func setupGinMiddleware(ginRouter *gin.Engine) {
	log.Info().Msg("Configuring GIN middleware")
	ginRouter.Use(gin.Recovery()) // Default recovery middleware
	ginRouter.Use(middleware.ZerologConsoleRequestLogging())
	ginRouter.Use(middleware.DisableCache())
	ginRouter.Use(middleware.AddSecurityHeaders(false))
}
