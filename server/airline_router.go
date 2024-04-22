package server

import (
	"airport-app-backend/controllers"
	"airport-app-backend/services"

	"gorm.io/gorm"
)

func (srv *AppServer) AirlineRouter(db *gorm.DB) {
	airlineService := services.NewServiceRepository(db)
	airlineController := controllers.NewAirlineControllerRepository(airlineService)
	
	srv.router.GET("/airline/", airlineController.HandleGetAirline)
	srv.router.GET("/airline/:id", airlineController.HandleGetAirlineById)
	srv.router.POST("/airline", airlineController.HandleCreateNewAirline)
}
