package server

import (
	"airport-app-backend/controllers"
	"airport-app-backend/services"

	"gorm.io/gorm"
)

func (srv *AppServer) AircraftRouter(db *gorm.DB) {
	service := services.NewServiceRepository(db)
	aircraftController := controllers.NewAircraftControllerRepository(service)

	srv.router.GET("/aircrafts", aircraftController.HandleGetAircrafts)
}
