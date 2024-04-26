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

func Problems(
	log *slog.Logger,
	problemProvider models.ProblemProvider,
	tagProvider models.TagProvider,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		difficulty := c.Query("difficulty")
		tagsNames := c.QueryArray("tags")

		tags, err := tagProvider.TagsByNames(c, tagsNames)
		if err != nil {
			log.Info("invalid request query parameters")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid form. Please, try again.")
			return
		}

		var problems []models.Problem
		if difficulty == "" && len(tags) == 0 {
			problems, err = problemProvider.Problems(c)
		} else {
			problems, err = problemProvider.ProblemsByFilters(c, &difficulty, tags)
		}

		if err != nil {
			log.Error("error retrieving the problem from db")
			templates.RespondWithError(c, http.StatusInternalServerError, "Internal server error. Sorry.")
			return
		}
		if problems == nil {
			log.Info("invalid problem filters provided")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid problem filters provided.")
			return
		}
		log.Info("getting problems list")
		templates.ProblemsResponse(c, problems, difficulty, tags)
	}
}

func ProblemByID(
	log *slog.Logger,
	problemProvider models.ProblemProvider,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Info("non-parsable problem id provided: " + c.Param("id"))
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid problem ID provided.")
			return
		}

		log = log.With(slog.Int("id", id))
		problem, err := problemProvider.Problem(c, uint(id))
		if err != nil {
			log.Error("error retrieving the problem from db")
			templates.RespondWithError(c, http.StatusInternalServerError, "Internal server error. Sorry.")
			return
		}
		if problem.ID == 0 {
			log.Info("invalid problem id provided")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid problem ID provided.")
			return
		}
		templates.ProblemResponse(c, problem)
	}
}

func CreateProblem(
	log *slog.Logger,
	problemSaver models.ProblemSaver,
	tagProvider models.TagProvider,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.PostForm("title")
		description := []byte(c.PostForm("description"))
		difficulty := c.PostForm("difficulty")
		tagsNames := strings.Split(c.PostForm("tagsNames"), " ")

		tags, err := tagProvider.TagsByNames(c, tagsNames)
		if err != nil {
			log.Info("invalid request form")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid form. Please, try again.")
			return
		}

		problem := models.ProblemRaw(title, description, difficulty, tags)
		id, err := problemSaver.SaveProblem(c, problem)
		if err != nil {
			log.Info("invalid request form")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid form. Please, try again.")
			return
		}

		log.Info(fmt.Sprintf("created a new problem with id %v", id))
		templates.CreateProblemResponse(c, problem)
	}
}

func NewProblemForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		templates.NewProblemFormResponse(c)
	}
}

func RemoveProblem(
	log *slog.Logger,
	problemRemover models.ProblemRemover,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Info("non-parsable problem id provided: " + c.Param("id"))
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid problem ID provided.")
			return
		}
		if err = problemRemover.RemoveProblem(c, uint(id)); err != nil {
			log.Info(fmt.Sprintf("invalid id provided: %v", id))
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid problem ID provided.")
			return
		}

		log.Info(fmt.Sprintf("removed problem with id %v successfully", id))
		templates.RemoveProblemResponse(c)
	}
}

func RemoveProblemForm(log *slog.Logger, problemProvider models.ProblemProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Info("non-parsable problem id provided: " + c.Param("id"))
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid problem ID provided.")
			return
		}
		problem, err := problemProvider.Problem(c, uint(id))
		if err != nil {
			log.Info("error retrieving problem from db")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid problem ID provided.")
			return
		}
		if problem.ID == 0 {
			log.Info("invalid problem id provided")
			templates.RespondWithError(c, http.StatusBadRequest, "Invalid problem ID provided.")
			return
		}
		templates.RemoveProblemFormResponse(c, problem)
	}
}
