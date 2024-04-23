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
	setupRoutes(engine, log, storage, storage, storage, storage, storage)
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
	tagProvider models.TagProvider,

) {
	r.GET("/", handlers.Index)
	r.GET("/article/:id", handlers.ArticleByID(log, articleProvider))
	r.GET("/new_article/", handlers.NewArticleForm(log))
	r.POST("/article/:id", handlers.CreateArticle(log, articleSaver, tagProvider))
}

func setupMiddleware(r *gin.Engine, log *slog.Logger) {
	r.Use(middleware.Logger(log))
	r.Use(gin.Recovery())
}
