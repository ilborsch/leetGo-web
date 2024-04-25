package templates

import (
	"github.com/gin-gonic/gin"
	"github.com/ilborsch/leetGo-web/internal/models"
	"net/http"
)

const (
	NewTagFormTemplateName = "new_tag_form.html"
	CreateTagTemplateName  = "create_tag.html"
	RemoveTagTemplateName  = "remove_tag.html"
)

func NewTagFormResponse(c *gin.Context) {
	c.HTML(http.StatusOK, NewTagFormTemplateName, gin.H{})
}

func CreateTagResponse(c *gin.Context, tag models.Tag) {
	c.HTML(http.StatusOK, CreateTagTemplateName, tag)
}

func RemoveTagResponse(c *gin.Context) {
	c.HTML(http.StatusOK, RemoveTagTemplateName, gin.H{})
}
