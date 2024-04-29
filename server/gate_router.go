package server

import (
	"airport-app-backend/controllers"
	"airport-app-backend/repositories"

	"gorm.io/gorm"
)

func (srv *AppServer) GateRouter(db *gorm.DB) {
	gateRepository := repositories.NewServiceRepository(db)
	gateController := controllers.NewGateRepository(gateRepository)

	srv.router.GET("/gates", gateController.HandleGetAllGates)
	srv.router.GET("/gate/:id", gateController.HandleGetGate)
	srv.router.POST("/gate", gateController.HandleCreateNewGate)
	srv.router.PUT("/gate/:id", gateController.HandleUpdateGate)
}
