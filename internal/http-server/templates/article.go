package templates

import (
	"github.com/gin-gonic/gin"
	"github.com/ilborsch/leetGo-web/internal/models"
	"net/http"
)

const (
	GetArticleTemplateName        = "article.html"
	CreateArticleTemplateName     = "create_article.html"
	NewArticleFormTemplateName    = "new_article_form.html"
	UpdateArticleTemplateName     = "update_article.html"
	RemoveArticleTemplateName     = "remove_article.html"
	UpdateArticleFormTemplateName = "update_article_form.html"
	RemoveArticleFormTemplateName = "remove_article_form.html"
)

func ArticleResponse(c *gin.Context, article models.Article) {
	c.HTML(http.StatusOK, GetArticleTemplateName, article)
}

func CreateArticleResponse(c *gin.Context, article models.Article) {
	c.HTML(http.StatusOK, CreateArticleTemplateName, article)
}

func NewArticleFormResponse(c *gin.Context) {
	c.HTML(http.StatusOK, NewArticleFormTemplateName, gin.H{})
}

func UpdateArticleFormResponse(c *gin.Context, article models.Article) {
	c.HTML(http.StatusOK, UpdateArticleFormTemplateName, article)
}

func UpdateArticleResponse(c *gin.Context, article models.Article) {
	c.HTML(http.StatusOK, UpdateArticleTemplateName, article)
}

func RemoveArticleResponse(c *gin.Context) {
	c.HTML(http.StatusOK, RemoveArticleTemplateName, gin.H{})
}

func RemoveArticleFormResponse(c *gin.Context, article models.Article) {
	c.HTML(http.StatusOK, RemoveArticleFormTemplateName, article)
}
