package server

import (
  "airport-app-backend/config"
  "airport-app-backend/database"
  "airport-app-backend/middleware"

  "github.com/gin-gonic/gin"
  "github.com/rs/zerolog/log"
)

// Main API routes
func (srv *AppServer) setupRoutesAndMiddleware() {

	log.Info().Msg("Connecting to postgres database")

	DB, err := database.ConnectToDB()
	if err != nil {
		log.Info().Err(err).Msg("Database connection failed")
		return
	}

	srv.router.Use(middleware.ZerologConsoleRequestLogging())

	srv.HealthRouter(DB)
	srv.GateRouter(DB)
	srv.AirlineRouter(DB)

	// Middleware
	log.Info().Msg("Configuring GIN middleware")
	srv.router.Use(gin.Recovery()) // Default recovery middleware

  srv.router.Use(middleware.DisableCache())
  srv.router.Use(middleware.AddSecurityHeaders(config.EnableTls))
  srv.router.Use(middleware.HandleFaviconRequests())

}
