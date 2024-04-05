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

		ctx.Next()

		latency := time.Since(startTime)

		logMsg := "Request"
		if len(ctx.Errors) > 0 {
			logMsg = ctx.Errors.String()
		}

		status := ctx.Writer.Status()
		requestPath := ctx.Request.URL.Path + "?" + ctx.Request.URL.RawQuery

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
