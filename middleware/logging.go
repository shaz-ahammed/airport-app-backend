package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"time"
)

var consoleLogger = zerolog.New(os.Stderr).Output(zerolog.ConsoleWriter{
	Out:        os.Stderr,
	NoColor:    false,
	TimeFormat: time.RFC1123Z,
})

var detailedLogging = os.Getenv("DETAILED_LOGGING")

func ZerologConsoleRequestLogging() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()

		var requestPath string

		if requestRawQuery := context.Request.URL.RawQuery; requestRawQuery == "" {
			requestPath = context.Request.URL.Path
		} else {
			requestPath = fmt.Sprintf("%s?%s", context.Request.URL.Path, requestRawQuery)
		}

		context.Next()

		latency := time.Since(startTime)

		subLogger := consoleLogger.With().Timestamp().
			Int("http-status", context.Writer.Status()).
			Str("method", context.Request.Method).
			Str("request-path", requestPath).
			Str("client-ip", context.ClientIP()).
			Str("user-agent", context.Request.UserAgent()).
			Int64("content-length", context.Request.ContentLength).
			Dur("latency", latency).
			Logger()

		logMsg := "Request"
		if len(context.Errors) > 0 {
			logMsg = context.Errors.String()
		}

		if context.Writer.Status() >= http.StatusInternalServerError {
			subLogger.Error().Msg(logMsg)
		} else if context.Writer.Status() >= http.StatusBadRequest && context.Writer.Status() < http.StatusInternalServerError {
			subLogger.Warn().Msg(logMsg)
		} else {
			if detailedLogging != "false" {
				subLogger.Info().Msg(logMsg)
			}
		}
	}
}
