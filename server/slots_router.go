package server

import (
	"airport-app-backend/controllers"
	"airport-app-backend/repositories"

	"gorm.io/gorm"
)

func (srv *AppServer) SlotRouter(db *gorm.DB) {
	slotRepository := repositories.NewServiceRepository(db)
	slotController := controllers.NewSlotController(slotRepository)

	srv.router.GET("/slots", slotController.HandleGetAllSlots)
}
