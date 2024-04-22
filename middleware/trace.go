package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"
)

func TraceSpanTags() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, span := trace.StartSpan(c.Request.Context(), c.FullPath())
		defer span.End()

		span.AddAttributes(
			trace.StringAttribute("http.method", c.Request.Method),
			trace.StringAttribute("http.url", c.Request.URL.Path),
			trace.StringAttribute("http.status_code", strconv.Itoa(c.Writer.Status())),
			trace.StringAttribute("trace.id", span.SpanContext().TraceID.String()),
			trace.StringAttribute("http.client_ip", c.ClientIP()),
			trace.StringAttribute("http.user_agent", c.Request.UserAgent()),
			trace.StringAttribute("http.content_length", strconv.FormatInt(c.Request.ContentLength, 10)),
		)
		c.Next()

	}
}
