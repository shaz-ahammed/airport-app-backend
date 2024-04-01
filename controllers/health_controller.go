package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (repo *ControllerRepository) HandleHealth(ctx *gin.Context) {
	log.Debug().Msg("Getting application health information")
	appHealth := repo.service.GetAppHealth()
	ctx.JSON(http.StatusOK, appHealth)
}
