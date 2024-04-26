package server

import (
	"airport-app-backend/controllers"
	"airport-app-backend/repositories"

	"gorm.io/gorm"
)

func (srv *AppServer) GateRouter(db *gorm.DB) {
	gateService := repositories.NewServiceRepository(db)
	gateController := controllers.NewGateRepository(gateService)

	srv.router.GET("/gates", gateController.HandleGetGates)
	srv.router.GET("/gate/:id", gateController.HandleGetGateById)
	srv.router.POST("/gate", gateController.HandleCreateNewGate)
	srv.router.PUT("/gate/:id", gateController.HandleUpdateGate)
}
