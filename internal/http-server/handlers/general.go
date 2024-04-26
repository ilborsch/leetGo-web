package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilborsch/leetGo-web/internal/http-server/templates"
	"log/slog"
)

func Home(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		templates.HomePageResponse(c)
	}
}
