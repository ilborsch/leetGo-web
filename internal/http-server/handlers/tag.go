package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilborsch/leetGo-web/internal/models"
	"log/slog"
)

func NewTagForm() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateTag(log *slog.Logger, tagSaver models.TagSaver) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func RemoveTag(log *slog.Logger, tagRemover models.TagRemover) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
