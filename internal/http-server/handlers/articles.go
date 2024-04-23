package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilborsch/leetGo-web/internal/http-server/templates"
	"github.com/ilborsch/leetGo-web/internal/models"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

func ArticleByID(
	log *slog.Logger,
	articleProvider models.ArticleProvider,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Info("non-parsable article id provided: " + c.Param("id"))
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid article ID provided.")
		}

		log = log.With(slog.Int("id", id))
		article, err := articleProvider.Article(c, uint(id))
		if err != nil {
			log.Error("error retrieving the article from db")
			templates.RespondWithError(c, http.StatusInternalServerError, "Internal server error. Sorry.")
		}
		if article.ID == 0 {
			log.Info("invalid article id provided")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid article ID provided.")
		}
		templates.ArticleResponse(c, article)
	}
}

func CreateArticle(
	log *slog.Logger,
	articleSaver models.ArticleSaver,
	tagProvider models.TagProvider,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		// implement middleware to check if auth token is valid
		// authorID, _ := c.Get("userID")
		authorID := uint(1)
		title := c.PostForm("title")
		content := []byte(c.PostForm("content"))
		_, isPublished := c.GetPostForm("isPublished")
		tagsNames := strings.Split(c.PostForm("tagsNames"), " ")

		tags := make([]models.Tag, 0, len(tagsNames))
		for _, tagName := range tagsNames {
			tag, err := tagProvider.TagByName(c, tagName)
			if err == nil {
				tags = append(tags, tag)
			}
		}

		article := models.ArticleRaw(title, content, authorID, isPublished, tags)
		id, err := articleSaver.SaveArticle(c, article)
		if err != nil {
			log.Info("invalid request form")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid form. Please, try again.")
		}
		log.Info(fmt.Sprintf("created a new article with id %v", id))
		templates.ArticleSuccessResponse(c, article)
	}
}

func NewArticleForm(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		templates.RespondWithNewArticleForm(c)
	}
}
