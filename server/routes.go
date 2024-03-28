package server

import (
	"airport-app-backend/config"
	"airport-app-backend/middleware"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Main API routes
func (srv *AppServer) setupRoutesAndMiddleware() {
	srv.router.GET("/health/", srv.handleHealth)
	srv.router.GET("/ping", srv.handlePing)
	srv.router.GET("/", srv.handleIndex)

	// Middleware
	log.Info().Msg("Configuring GIN middleware")
	srv.router.Use(gin.Recovery()) // Default recovery middleware

	srv.router.Use(middleware.DisableCache())
	srv.router.Use(middleware.AddSecurityHeaders(config.EnableTls))
	srv.router.Use(middleware.HandleFaviconRequests())

	if config.EnableDetailedRequestLogging {
		log.Info().Msg("Enabling request logging middleware")
		srv.router.Use(middleware.ZerologConsoleRequestLogging())
	}
}
