package main

import (
	"github.com/ilborsch/leetGo-web/internal/config"
	"github.com/ilborsch/leetGo-web/internal/http-server/router"
	"github.com/ilborsch/leetGo-web/internal/logger"
	"github.com/ilborsch/leetGo-web/internal/storage"
)

func main() {
	// TODO: Frontend
	// TODO: blocks in html to reuse code
	// TODO: sso grpc client
	// TODO: Implement users
	// TODO: execution-engine client
	// TODO: Selenium testing

	// init config
	cfg := config.MustLoad()

	// init logger
	log := logger.SetupLogger(cfg.Env)
	log.Info("starting application")

	// init database
	db := storage.New()

	// init router
	r := router.New(log, db)

	r.Run("0.0.0.0", cfg.Port)
}
