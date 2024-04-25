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
			return
		}

		log = log.With(slog.Int("id", id))
		article, err := articleProvider.Article(c, uint(id))
		if err != nil {
			log.Error("error retrieving the article from db")
			templates.RespondWithError(c, http.StatusInternalServerError, "Internal server error. Sorry.")
			return
		}
		if article.ID == 0 {
			log.Info("invalid article id provided")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid article ID provided.")
			return
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
		isPublished := c.PostForm("isPublished") == "on"
		tagsNames := strings.Split(c.PostForm("tagsNames"), ", ")

		tags, err := tagProvider.TagsByNames(c, tagsNames)
		if err != nil {
			log.Info("invalid request form")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid form. Please, try again.")
			return
		}

		article := models.ArticleRaw(title, content, authorID, isPublished, tags)
		id, err := articleSaver.SaveArticle(c, article)
		if err != nil {
			log.Info("invalid request form")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid form. Please, try again.")
			return
		}

		log.Info(fmt.Sprintf("created a new article with id %v", id))
		templates.CreateArticleResponse(c, article)
	}
}

func NewArticleForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		templates.NewArticleFormResponse(c)
	}
}

func UpdateArticleForm(log *slog.Logger, articleProvider models.ArticleProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Info("non-parsable article id provided: " + c.Param("id"))
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid article ID provided.")
			return
		}
		article, err := articleProvider.Article(c, uint(id))
		if err != nil {
			log.Info("error retrieving article from db")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid article ID provided.")
			return
		}
		if article.ID == 0 {
			log.Info("invalid article id provided")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid article ID provided.")
			return
		}
		templates.UpdateArticleFormResponse(c, article)
	}
}

func RemoveArticleForm(log *slog.Logger, articleProvider models.ArticleProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Info("non-parsable article id provided: " + c.Param("id"))
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid article ID provided.")
			return
		}
		article, err := articleProvider.Article(c, uint(id))
		if err != nil {
			log.Info("error retrieving article from db")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid article ID provided.")
			return
		}
		if article.ID == 0 {
			log.Info("invalid article id provided")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid article ID provided.")
			return
		}
		templates.RemoveArticleFormResponse(c, article)
	}
}

func UpdateArticle(
	log *slog.Logger,
	articleUpdater models.ArticleUpdater,
	tagProvider models.TagProvider,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		// implement middleware to check if auth token is valid
		// authorID, _ := c.Get("userID")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Info("non-parsable article id provided: " + c.Param("id"))
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid article ID provided.")
			return
		}
		authorID := uint(1)
		title := c.PostForm("title")
		content := []byte(c.PostForm("content"))
		_, isPublished := c.GetPostForm("isPublished")
		tagsNames := strings.Split(c.PostForm("tagsNames"), ", ")

		tags, err := tagProvider.TagsByNames(c, tagsNames)
		if err != nil {
			log.Info("invalid request form")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid form. Please, try again.")
			return
		}

		article := models.ArticleRaw(title, content, authorID, isPublished, tags)
		err = articleUpdater.UpdateArticle(c, uint(id), article)
		if err != nil {
			log.Info("invalid request form")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid form. Please, try again.")
			return
		}

		log.Info(fmt.Sprintf("updated article with id %v successfully", id))
		templates.UpdateArticleResponse(c, article)
	}
}

func RemoveArticle(
	log *slog.Logger,
	articleRemover models.ArticleRemover,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Info("non-parsable article id provided: " + c.Param("id"))
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid article ID provided.")
			return
		}
		if err = articleRemover.RemoveArticle(c, uint(id)); err != nil {
			log.Info(fmt.Sprintf("invalid id provided: %v", id))
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid article ID provided.")
			return
		}

		log.Info(fmt.Sprintf("removed article with id %v successfully", id))
		templates.RemoveArticleResponse(c)
	}
}
