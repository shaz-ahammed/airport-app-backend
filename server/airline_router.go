package server

import (
	"airport-app-backend/controllers"
	"airport-app-backend/repositories"

	"gorm.io/gorm"
)

func (srv *AppServer) AirlineRouter(db *gorm.DB) {
	airlineRepository := repositories.NewServiceRepository(db)
	airlineController := controllers.NewAirlineControllerRepository(airlineRepository)

	srv.router.GET("/airlines", airlineController.HandleGetAllAirlines)
	srv.router.GET("/airline/:id", airlineController.HandleGetAirlineById)
	srv.router.POST("/airline", airlineController.HandleCreateNewAirline)
	srv.router.DELETE("/airline/:id", airlineController.HandleDeleteAirlineById)
}
