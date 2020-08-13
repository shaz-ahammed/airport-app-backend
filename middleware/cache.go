package middleware

import "github.com/gin-gonic/gin"

// Middleware to add cache control headers to server responses disabling client cache.
func DisableCache() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Add("Expires", "0")
		ctx.Writer.Header().Add("Pragma", "no-cache")
		// For Cache-Control header only no-store directive is needed
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control
		ctx.Writer.Header().Add("Cache-Control", "no-store")
		ctx.Next()
	}
}
