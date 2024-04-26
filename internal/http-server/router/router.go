package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilborsch/leetGo-web/internal/http-server/handlers"
	"github.com/ilborsch/leetGo-web/internal/http-server/middleware"
	"github.com/ilborsch/leetGo-web/internal/models"
	"github.com/ilborsch/leetGo-web/internal/storage"
	"html/template"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type Router struct {
	log    *slog.Logger
	engine *gin.Engine
}

func (r *Router) Run(address string, port int) {
	addr := fmt.Sprintf("%s:%v", address, port)
	if err := r.engine.Run(addr); err != nil {
		panic("error starting gin engine")
	}
}

func New(log *slog.Logger, storage *storage.Storage) *Router {
	engine := gin.New()
	setupHTMLTemplates(engine)
	setupMiddleware(engine, log)
	// I am very sorry about that :)
	setupRoutes(engine, log, storage, storage, storage, storage, storage, storage, storage, storage, storage, storage)
	return &Router{
		log:    log,
		engine: engine,
	}
}

func setupHTMLTemplates(engine *gin.Engine) {
	const directoryName = "static/html"
	tmpl := template.New("")
	files, err := os.ReadDir(directoryName)
	if err != nil {
		panic(fmt.Sprintf("couldn't read template files from ./%s directory", directoryName))
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".html") {
			continue
		}
		filePath := filepath.Join(directoryName, file.Name())
		_, err := tmpl.ParseFiles(filePath)
		if err != nil {
			panic(fmt.Sprintf("couldn't parse template file ./%s", filePath))
		}
	}
	engine.SetHTMLTemplate(tmpl)
}

func setupMiddleware(r *gin.Engine, log *slog.Logger) {
	r.Use(middleware.Logger(log))
	r.Use(gin.Recovery())
}

func setupRoutes(
	r *gin.Engine,
	log *slog.Logger,
	articleProvider models.ArticleProvider,
	articleSaver models.ArticleSaver,
	articleUpdater models.ArticleUpdater,
	articleRemover models.ArticleRemover,
	problemProvider models.ProblemProvider,
	problemSaver models.ProblemSaver,
	problemRemover models.ProblemRemover,
	tagProvider models.TagProvider,
	tagSaver models.TagSaver,
	tagRemover models.TagRemover,
) {
	r.GET("/", handlers.Home(log))

	articleGroup := r.Group("/articles")
	articleGroup.GET("/:id", handlers.ArticleByID(log, articleProvider))
	articleGroup.POST("/new", handlers.CreateArticle(log, articleSaver, tagProvider))
	articleGroup.PATCH("/:id", handlers.UpdateArticle(log, articleUpdater, tagProvider))
	articleGroup.DELETE("/:id", handlers.RemoveArticle(log, articleRemover))
	articleGroup.GET("/new", handlers.NewArticleForm())
	articleGroup.GET("/update/:id", handlers.UpdateArticleForm(log, articleProvider))
	articleGroup.GET("/remove/:id", handlers.RemoveArticleForm(log, articleProvider))

	problemGroup := r.Group("/problems")
	problemGroup.GET("/:id", handlers.ProblemByID(log, problemProvider))
	problemGroup.GET("/", handlers.ProblemsList(log, problemProvider, tagProvider))
	problemGroup.GET("/new", handlers.NewProblemForm())
	problemGroup.POST("/new", handlers.CreateProblem(log, problemSaver, tagProvider))
	problemGroup.GET("/remove/:id", handlers.RemoveProblemForm(log, problemProvider))
	problemGroup.DELETE("/:id", handlers.RemoveProblem(log, problemRemover))

	tagGroup := r.Group("/tags")
	tagGroup.GET("/new", handlers.NewTagForm())
	tagGroup.POST("/new", handlers.CreateTag(log, tagSaver))
	tagGroup.GET("/remove/:id", handlers.RemoveTagForm(log, tagProvider))
	tagGroup.DELETE("/:id", handlers.RemoveTag(log, tagRemover))
}
