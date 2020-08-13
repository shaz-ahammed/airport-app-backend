package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (srv *AppServer) handlePing(ctx *gin.Context) {
	log.Debug().Msg("Replying to Ping")
	ctx.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
}
