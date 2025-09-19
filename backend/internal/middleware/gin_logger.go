package middleware

import (
	"backend/internal/pkg/logger"
	"backend/internal/response"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)

		log := logger.WithRequest(c)

		event := log.Info()
		if c.Writer.Status() >= 400 {
			event = log.Error()
		}

		event.
			Str("method", c.Request.Method).
			Int("status", c.Writer.Status()).
			Str("path", c.Request.URL.Path).
			Str("query", c.Request.URL.RawQuery).
			Int("body_size", c.Writer.Size()).
			Str("latency", latency.String()).
			Str("client_ip", c.ClientIP()).
			Str("user_agent", c.Request.UserAgent()).
			Msg("HTTP request")
	}
}

func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log := logger.WithRequest(c)
				log.Error().
					Interface("error", err).
					Str("stack", string(debug.Stack())).
					Msg("HTTP panic recovered")

				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					response.NewErrorMessage("Internal server error"),
				)
			}
		}()

		c.Next()
	}
}
