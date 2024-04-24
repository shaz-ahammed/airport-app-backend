package server

import (
	"airport-app-backend/controllers"
	"airport-app-backend/repositories"

	"gorm.io/gorm"
)

func (srv *AppServer) AirlineRouter(db *gorm.DB) {
	airlineRepository := repositories.NewServiceRepository(db)
	airlineController := controllers.NewAirlineController(airlineRepository)

	srv.router.GET("/airlines", airlineController.HandleGetAllAirlines)
	srv.router.GET("/airline/:id", airlineController.HandleGetAirline)
	srv.router.POST("/airline", airlineController.HandleCreateNewAirline)
	srv.router.PUT("/airline/:id", airlineController.HandleUpdateAirline)
	srv.router.DELETE("/airline/:id", airlineController.HandleDeleteAirline)
}
