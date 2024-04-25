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
	engine.SetHTMLTemplate(template.Must(template.ParseFiles("static/html/*.html")))

	setupMiddleware(engine, log)
	// i am very sorry about that :)
	setupRoutes(engine, log, storage, storage, storage, storage, storage, storage, storage, storage, storage, storage)
	return &Router{
		log:    log,
		engine: engine,
	}
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
	r.GET("/", handlers.Index)

	articleGroup := r.Group("/article")
	articleGroup.GET("/:id", handlers.ArticleByID(log, articleProvider))
	articleGroup.GET("/new", handlers.NewArticleForm())
	articleGroup.POST("/new", handlers.CreateArticle(log, articleSaver, tagProvider))
	articleGroup.PATCH("/:id", handlers.UpdateArticle(log, articleUpdater, tagProvider))
	articleGroup.DELETE("/:id", handlers.RemoveArticle(log, articleRemover))

	problemGroup := r.Group("/problem")
	problemGroup.GET("/:id", handlers.ProblemByID(log, problemProvider))
	problemGroup.GET("/", handlers.ProblemsList(log, problemProvider, tagProvider))
	problemGroup.GET("/new", handlers.NewProblemForm())
	problemGroup.POST("/new", handlers.CreateProblem(log, problemSaver, tagProvider))
	problemGroup.DELETE("/:id", handlers.RemoveProblem(log, problemRemover))

	tagGroup := r.Group("/problem")
	tagGroup.GET("/new", handlers.NewTagForm())
	tagGroup.POST("/new", handlers.CreateTag(log, tagSaver))
	tagGroup.DELETE("/:id", handlers.RemoveTag(log, tagRemover))
}

func setupMiddleware(r *gin.Engine, log *slog.Logger) {
	r.Use(middleware.Logger(log))
	r.Use(gin.Recovery())
}
