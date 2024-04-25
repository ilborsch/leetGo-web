package middleware

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

func Logger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		method := c.Request.Method
		url := c.Request.URL.String()
		statusCode := c.Writer.Status()
		log := log.With(
			slog.String("method", method),
			slog.String("path", url),
			slog.Int("statusCode", statusCode),
		)
		log.Info("request")
	}
}
