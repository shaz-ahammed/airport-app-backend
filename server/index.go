package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (srv *AppServer) handleIndex(ctx *gin.Context) {
	log.Debug().Msg("Handling index")
	ctx.String(http.StatusOK, "Hello World "+fmt.Sprint(time.Now().Unix()))
}
