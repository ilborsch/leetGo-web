package templates

import (
	"github.com/gin-gonic/gin"
)

const ErrorHTMLTemplateName = "error.html"

func RespondWithError(c *gin.Context, code int, error string) {
	c.HTML(code, ErrorHTMLTemplateName, gin.H{
		"Code":  code,
		"Error": error,
	})
}
