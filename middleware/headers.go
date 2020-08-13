package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Middleware to add security headers to server responses.
// When isTlsEnabled is set to true HSTS header will be added and 'upgrade-insecure-requests' will be added to CSP header.
func AddSecurityHeaders(isTlsEnabled bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Add("X-Frame-Options", "DENY")
		ctx.Writer.Header().Add("X-Content-Type-Options", "nosniff")
		ctx.Writer.Header().Add("Referrer-Policy", "no-referrer")

		// HSTS - HTTP Strict Transport Security
		if isTlsEnabled {
			hstsMaxAge := 60 * 60 * 24 * 365 * 3 // 3 years
			hstsValue := fmt.Sprintf("max-age=%d ;includeSubDomains; preload", hstsMaxAge)
			ctx.Writer.Header().Add("Strict-Transport-Security", hstsValue)
		}

		// CSP - Content-Security-Policy
		cspPolicy := "default-src 'none';"
		if isTlsEnabled {
			cspPolicy += " upgrade-insecure-requests;"
		}
		ctx.Writer.Header().Add("Content-Security-Policy", cspPolicy)

		ctx.Next()
	}
}
