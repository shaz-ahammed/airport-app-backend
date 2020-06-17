package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddSecurityHeaders(isTlsEnabled bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Add("X-Frame-Options", "DENY")
		context.Writer.Header().Add("X-Content-Type-Options", "nosniff")
		context.Writer.Header().Add("Content-Security-Policy", "default-src 'none';")
		context.Writer.Header().Add("Referrer-Policy", "no-referrer")

		if isTlsEnabled {
			hstsMaxAge := 60 * 60 * 24 * 365 * 3 // 3 years
			hstsValue := fmt.Sprintf("max-age=%d ;includeSubDomains; preload", hstsMaxAge)
			context.Writer.Header().Add("Strict-Transport-Security", hstsValue)
		}

		context.Next()
	}
}
