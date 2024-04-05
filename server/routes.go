package server

import (
	"airport-app-backend/config"
	"airport-app-backend/middleware"
	"airport-app-backend/services"

	"airport-app-backend/controllers"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Main API routes
func (srv *AppServer) setupRoutesAndMiddleware() {

	log.Info().Msg("Connecting to postgres database")

	DB, err := ConnectToDB()
	if err != nil {
		log.Info().Err(err).Msg("Database connection failed")
		return
	}

	err = MigrateAll(DB)
	if err != nil {
		log.Info().Err(err).Msg("Database migration failed")
		return
	}
	log.Info().Msg("Database migration Successful")

	serviceRepo := services.NewServiceRepository(DB)
	controllerRepo := controllers.NewControllerRepository(serviceRepo)

	srv.router.Use(middleware.ZerologConsoleRequestLogging())

	srv.router.GET("/health/", controllerRepo.HandleHealth)

	// Middleware
	log.Info().Msg("Configuring GIN middleware")
	srv.router.Use(gin.Recovery()) // Default recovery middleware

	srv.router.Use(middleware.DisableCache())
	srv.router.Use(middleware.AddSecurityHeaders(config.EnableTls))
	srv.router.Use(middleware.HandleFaviconRequests())

}
