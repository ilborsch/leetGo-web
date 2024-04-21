package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilborsch/leetGo-web/internal/config"
	"github.com/ilborsch/leetGo-web/internal/handlers"
	"github.com/ilborsch/leetGo-web/internal/logger"
	"github.com/ilborsch/leetGo-web/internal/middleware"
	"log/slog"
)

func main() {
	// init config
	cfg := config.MustLoad()

	// init logger
	log := logger.SetupLogger(cfg.Env)
	log.Info("starting application")

	// init router
	r := gin.New()

	// setup routes and middleware
	setupMiddleware(r, log)
	setupRoutes(r)

	// start app
	addr := fmt.Sprintf("0.0.0.0:%v", cfg.Port)
	if err := r.Run(addr); err != nil {
		fmt.Println("fatal error while run")
		return
	}
}

func setupRoutes(r *gin.Engine) {
	r.GET("/", handlers.Index)
}

func setupMiddleware(r *gin.Engine, log *slog.Logger) {
	r.Use(middleware.Logger(log))
}
