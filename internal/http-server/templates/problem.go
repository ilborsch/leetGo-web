package templates

import (
	"github.com/gin-gonic/gin"
	"github.com/ilborsch/leetGo-web/internal/models"
	"net/http"
)

const (
	GetProblemTemplateName        = "problem.html"
	GetProblemsTemplateName       = "problems.html"
	CreateProblemTemplateName     = "create_problem.html"
	NewProblemFormTemplateName    = "new_problem_form.html"
	RemoveProblemTemplateName     = "remove_problem.html"
	RemoveProblemFormTemplateName = "remove_problem_form.html"
)

func ProblemsResponse(c *gin.Context, problems []models.Problem, difficulty string, tags []models.Tag) {
	c.HTML(http.StatusOK, GetProblemsTemplateName, gin.H{
		"Problems":   problems,
		"Difficulty": difficulty,
		"Tags":       tags,
	})
}

func ProblemResponse(c *gin.Context, problem models.Problem) {
	c.HTML(http.StatusOK, GetProblemTemplateName, problem)
}

func CreateProblemResponse(c *gin.Context, problem models.Problem) {
	c.HTML(http.StatusOK, CreateProblemTemplateName, problem)
}

func NewProblemFormResponse(c *gin.Context) {
	c.HTML(http.StatusOK, NewProblemFormTemplateName, gin.H{})
}

func RemoveProblemResponse(c *gin.Context) {
	c.HTML(http.StatusOK, RemoveProblemTemplateName, gin.H{})
}

func RemoveProblemFormResponse(c *gin.Context, problem models.Problem) {
	c.HTML(http.StatusOK, RemoveProblemFormTemplateName, problem)
}
