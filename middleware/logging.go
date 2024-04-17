package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func setupLogger() zerolog.Logger {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z}
	return zerolog.New(consoleWriter).With().Timestamp().Logger()
}

var logger = setupLogger()

func ZerologConsoleRequestLogging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		latency := time.Since(startTime)

		logMsg := "Request"
		if len(ctx.Errors) > 0 {
			logMsg = ctx.Errors.String()
		}

		status := ctx.Writer.Status()
		var requestPath string

		if requestRawQuery := ctx.Request.URL.RawQuery; requestRawQuery == "" {
			requestPath = ctx.Request.URL.Path
		} else {
			requestPath = fmt.Sprintf("%s?%s", ctx.Request.URL.Path, requestRawQuery)
		}

		subLogger := logger.With().
			Int("http-status", ctx.Writer.Status()).
			Str("method", ctx.Request.Method).
			Str("request-path", requestPath).
			Str("client-ip", ctx.ClientIP()).
			Str("user-agent", ctx.Request.UserAgent()).
			Dur("latency", latency).
			Logger()

		if status >= http.StatusInternalServerError {
			subLogger.Error().Msg(logMsg)
		} else if status >= http.StatusBadRequest && status < http.StatusInternalServerError {
			subLogger.Warn().Msg(logMsg)
		} else {
			subLogger.Info().Msg(logMsg)
		}
	}
}
