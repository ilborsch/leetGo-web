package templates

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ErrorTemplateName = "error.html"
	HomePageTemplateName
)

func RespondWithError(c *gin.Context, code int, error string) {
	c.HTML(code, ErrorTemplateName, gin.H{
		"Code":  code,
		"Error": error,
	})
}

func HomePageResponse(c *gin.Context) {
	c.HTML(http.StatusOK, HomePageTemplateName, gin.H{})
}
