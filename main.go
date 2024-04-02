package main

import (
	"airport-app-backend/config"
	"airport-app-backend/server"

	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var Database *gorm.DB

func main() {
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
