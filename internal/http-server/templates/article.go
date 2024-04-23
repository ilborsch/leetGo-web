package templates

import (
	"github.com/gin-gonic/gin"
	"github.com/ilborsch/leetGo-web/internal/models"
	"net/http"
)

const (
	GetArticleHTMLTemplateName     = "article.html"
	ArticleSuccessHTMLTemplateName = "article_success.html"
)

func ArticleResponse(c *gin.Context, article models.Article) {
	c.HTML(http.StatusOK, GetArticleHTMLTemplateName, article)
}

func ArticleSuccessResponse(c *gin.Context, article models.Article) {
	c.HTML(http.StatusOK, ArticleSuccessHTMLTemplateName, article)
}

func RespondWithNewArticleForm(c *gin.Context) {

}
