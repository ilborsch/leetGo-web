package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilborsch/leetGo-web/internal/http-server/handlers"
	"github.com/ilborsch/leetGo-web/internal/http-server/middleware"
	"github.com/ilborsch/leetGo-web/internal/storage"
	"github.com/ilborsch/leetGo-web/pkg/sso"
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

func New(log *slog.Logger, storage *storage.Storage, ssoClient *sso.Client) *Router {
	engine := gin.New()
	setupHTMLTemplates(engine)
	setupMiddleware(engine, log)
	setupRoutes(engine, log, ssoClient, storage)
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
	ssoClient *sso.Client,
	storage *storage.Storage,
) {
	setupGeneralRoutes(r, log)
	setupArticleRoutes(r, log, storage)
	setupProblemRoutes(r, log, storage)
	setupTagRoutes(r, log, storage)
	setupUserRoutes(r, log, storage, ssoClient)
}

func setupGeneralRoutes(r *gin.Engine, log *slog.Logger) {
	r.GET("/", handlers.Home(log))

}

func setupArticleRoutes(r *gin.Engine, log *slog.Logger, storage *storage.Storage) {
	articleGroup := r.Group("/articles")
	articleGroup.GET("/", handlers.Articles(log, storage, storage))
	articleGroup.GET("/:id", handlers.ArticleByID(log, storage))
	articleGroup.POST("/new", handlers.CreateArticle(log, storage, storage))
	articleGroup.PATCH("/:id", handlers.UpdateArticle(log, storage, storage))
	articleGroup.DELETE("/:id", handlers.RemoveArticle(log, storage))
	articleGroup.GET("/new", handlers.NewArticleForm())
	articleGroup.GET("/update/:id", handlers.UpdateArticleForm(log, storage))
	articleGroup.GET("/remove/:id", handlers.RemoveArticleForm(log, storage))
}

func setupProblemRoutes(r *gin.Engine, log *slog.Logger, storage *storage.Storage) {
	problemGroup := r.Group("/problems")
	problemGroup.GET("/", handlers.Problems(log, storage, storage))
	problemGroup.GET("/:id", handlers.ProblemByID(log, storage))
	problemGroup.GET("/new", handlers.NewProblemForm())
	problemGroup.POST("/new", handlers.CreateProblem(log, storage, storage))
	problemGroup.GET("/remove/:id", handlers.RemoveProblemForm(log, storage))
	problemGroup.DELETE("/:id", handlers.RemoveProblem(log, storage))
}

func setupTagRoutes(r *gin.Engine, log *slog.Logger, storage *storage.Storage) {
	tagGroup := r.Group("/tags")
	tagGroup.GET("/new", handlers.NewTagForm())
	tagGroup.POST("/new", handlers.CreateTag(log, storage))
	tagGroup.GET("/remove/:id", handlers.RemoveTagForm(log, storage))
	tagGroup.DELETE("/:id", handlers.RemoveTag(log, storage))
}

func setupUserRoutes(r *gin.Engine, log *slog.Logger, storage *storage.Storage, ssoClient *sso.Client) {

}
