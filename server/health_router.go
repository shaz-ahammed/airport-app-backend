package server

import (
	"airport-app-backend/controllers"
	"airport-app-backend/services"

	"gorm.io/gorm"
)

func (srv *AppServer) HealthRouter(db *gorm.DB) {
	healthService := services.NewServiceRepository(db)
	healthController := controllers.NewControllerRepository(healthService)
	srv.router.GET("/health/", healthController.HandleHealth)

	srv.router.GET("/", healthController.Home)

}
