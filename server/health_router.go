package server

import (
	"airport-app-backend/controllers"
	"airport-app-backend/repositories"

	"gorm.io/gorm"
)

func (srv *AppServer) HealthRouter(db *gorm.DB) {
	healthRepository := repositories.NewServiceRepository(db)
	healthController := controllers.NewController(healthRepository)

	srv.router.GET("/health", healthController.HandleHealth)
	srv.router.GET("/", healthController.Home)
}
