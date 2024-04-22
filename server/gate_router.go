package server

import (
	"airport-app-backend/controllers"
	"airport-app-backend/services"

	"gorm.io/gorm"
)

func (srv *AppServer) GateRouter(db *gorm.DB) {
	gateService := services.NewServiceRepository(db)
	gateController := controllers.NewGateRepository(gateService)

	srv.router.GET("/gates", gateController.HandleGetGates)
	srv.router.GET("/gate/:id", gateController.HandleGetGateById)
}
