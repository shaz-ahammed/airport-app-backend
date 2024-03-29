package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"airport-app-backend/services"

)

func  HandleHealth(ctx *gin.Context) {
	log.Debug().Msg("Getting application health information")
	appHealth := services.GetAppHealth()
	ctx.JSON(http.StatusOK, appHealth)
}
