package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilborsch/leetGo-web/internal/http-server/templates"
	"github.com/ilborsch/leetGo-web/internal/models"
	"log/slog"
	"net/http"
	"strconv"
)

func NewTagForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		templates.NewTagFormResponse(c)
	}
}

func CreateTag(log *slog.Logger, tagSaver models.TagSaver) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		tag := models.TagRaw(name)
		id, err := tagSaver.SaveTag(c, tag)
		if err != nil {
			log.Error("error creating a new tag object")
			templates.RespondWithError(c, http.StatusInternalServerError, "Internal server error. Sorry.")
			return
		}
		log.With(slog.Int("id", int(id))).Info("created a new tag")
		templates.CreateTagResponse(c, tag)
	}
}

func RemoveTag(log *slog.Logger, tagRemover models.TagRemover) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Info("non-parsable tag id provided")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid tag id provided.")
			return
		}
		log = log.With(slog.Int("id", id))
		if err = tagRemover.RemoveTag(c, uint(id)); err != nil {
			log.Info("error removing the tag object")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid tag id provided.")
			return
		}
		log.Info("removed tag successfully")
		templates.RemoveTagResponse(c)
	}
}

func RemoveTagForm(log *slog.Logger, tagProvider models.TagProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Info("non-parsable tag id provided: " + c.Param("id"))
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid tag ID provided.")
			return
		}
		tag, err := tagProvider.Tag(c, id)
		if err != nil {
			log.Info("error retrieving tag from db")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid tag ID provided.")
			return
		}
		if tag.ID == 0 {
			log.Info("invalid tag id provided")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid tag ID provided.")
			return
		}
		templates.RemoveTagFormResponse(c, tag)
	}
}
