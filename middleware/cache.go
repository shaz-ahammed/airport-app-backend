package middleware

import "github.com/gin-gonic/gin"

func DisableCache() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Add("Expires", "0")
		context.Writer.Header().Add("Pragma", "no-cache")
		// For Cache-Control header only no-store directive is needed
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control
		context.Writer.Header().Add("Cache-Control", "no-store")
		context.Next()
	}
}
