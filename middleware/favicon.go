package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Middleware to handle requests for nonexistent favicon.ico. Returns HTTP 204 No Content
func HandleFaviconRequests() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/favicon.ico" {
			log.Info().Msg("Returning HTTP 204 No Content for favicon.ico request")
			ctx.AbortWithStatus(http.StatusNoContent)
		}
	}
}
