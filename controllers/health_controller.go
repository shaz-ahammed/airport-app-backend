package controllers

import (
	"fmt"
	"net/http"

	"airport-app-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (r *Repository) HandleHealth(ctx *gin.Context) {
	log.Debug().Msg("Getting application health information")
	appHealth := services.GetAppHealth()
	fmt.Println(r.db)
	ctx.JSON(http.StatusOK, appHealth)
}
