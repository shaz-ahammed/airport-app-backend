package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var consoleLogger = zerolog.New(os.Stderr).Output(zerolog.ConsoleWriter{
	Out:        os.Stderr,
	NoColor:    false,
	TimeFormat: time.RFC1123Z,
})

// Middleware that adds detailed request logging using rs/zerolog.
func ZerologConsoleRequestLogging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		var requestPath string

		if requestRawQuery := ctx.Request.URL.RawQuery; requestRawQuery == "" {
			requestPath = ctx.Request.URL.Path
		} else {
			requestPath = fmt.Sprintf("%s?%s", ctx.Request.URL.Path, requestRawQuery)
		}

		ctx.Next()

		latency := time.Since(startTime)

		subLogger := consoleLogger.With().Timestamp().
			Int("http-status", ctx.Writer.Status()).
			Str("method", ctx.Request.Method).
			Str("request-path", requestPath).
			Str("client-ip", ctx.ClientIP()).
			Str("user-agent", ctx.Request.UserAgent()).
			Int64("content-length", ctx.Request.ContentLength).
			Dur("latency", latency).
			Logger()

		logMsg := "Request"
		if len(ctx.Errors) > 0 {
			logMsg = ctx.Errors.String()
		}

		if ctx.Writer.Status() >= http.StatusInternalServerError {
			subLogger.Error().Msg(logMsg)
		} else if ctx.Writer.Status() >= http.StatusBadRequest && ctx.Writer.Status() < http.StatusInternalServerError {
			subLogger.Warn().Msg(logMsg)
		} else {
			subLogger.Info().Msg(logMsg)
		}
	}
}
