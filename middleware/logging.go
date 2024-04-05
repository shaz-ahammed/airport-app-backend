package middleware

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// Setup logging to both console and file.
func setupLogger() zerolog.Logger {
	logFile, err := os.OpenFile("requests.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open log file: %v", err)
		os.Exit(1)
	}

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z}
	multiWriter := io.MultiWriter(consoleWriter, logFile)

	return zerolog.New(multiWriter).With().Timestamp().Logger()
}

var logger = setupLogger()

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
		} else if status >= http.StatusBadRequest && status < http.StatusInternalServerError {
			subLogger.Warn().Msg(logMsg)
		} else {
			subLogger.Info().Msg(logMsg)
		}
	}
}
