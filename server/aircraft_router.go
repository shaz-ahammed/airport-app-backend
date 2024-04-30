package server

import (
	"airport-app-backend/controllers"
	"airport-app-backend/repositories"

	"gorm.io/gorm"
)

func (srv *AppServer) AircraftRouter(db *gorm.DB) {
	aircraftRepository := repositories.NewServiceRepository(db)
	aircraftController := controllers.NewAircraftController(aircraftRepository)

	srv.router.GET("/aircrafts", aircraftController.HandleGetAllAircrafts)
	srv.router.GET("/aircraft/:id", aircraftController.HandleGetAircraft)
}
