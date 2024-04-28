package main

import (
	"github.com/ilborsch/leetGo-web/internal/config"
	"github.com/ilborsch/leetGo-web/internal/http-server/router"
	"github.com/ilborsch/leetGo-web/internal/logger"
	"github.com/ilborsch/leetGo-web/internal/storage"
	"github.com/ilborsch/leetGo-web/pkg/sso"
)

func main() {
	// init config
	cfg := config.MustLoad()

	// init logger
	log := logger.SetupLogger(cfg.Env)
	log.Info("starting application")

	// init database
	db := storage.New()

	// init sso grpc client
	ssoClient := sso.NewClient(log, "0.0.0.0", cfg.SSO.Port, cfg.SSO.AppID)

	// init router
	r := router.New(log, db, ssoClient)

	r.Run("0.0.0.0", cfg.Port)
}
