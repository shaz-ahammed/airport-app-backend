package controllers

import (
	"airport-app-backend/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var AVAILABLE_SLOTS = true

type SlotController struct {
	repository repositories.ISlotRepository
}

func NewSlotController(repository repositories.ISlotRepository) *SlotController {
	return &SlotController{
		repository: repository,
	}
}

// @Summary Get all Slots
// @Router /slots [get]
// @Description get all the slots
// @ID get-all-slots
// @Tags slot
// @Produce  json
// @Param   page        query    int     false        "Page number (default = 1)"
// @Param   is_available        query    boolean     false        "Availablity of the slot. [true, false] (default = true)"
// @Success 200  "ok"
// @Failure 500 "Internal server error"
func (sc SlotController) HandleGetAllSlots(ctx *gin.Context) {
	// TODO: Convert to using a pagination library to handle this and other edge cases
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page < 0 {
		ctx.JSON(400, gin.H{"msg": "Page number must be greater than 0"})
		return
	}
	status, err := strconv.ParseBool(ctx.Query("is_available"))
	if err != nil {
		status = AVAILABLE_SLOTS
	}

	slots, err := sc.repository.RetrieveAllSlots(page, status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, slots)
}
