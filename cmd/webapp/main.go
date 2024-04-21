package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilborsch/leetGo-web/internal/config"
	handlers "github.com/ilborsch/leetGo-web/internal/handlers/index"
)

func main() {
	// init config
	cfg := config.MustLoad()
	fmt.Println(*cfg)

	// init logger

	// init gin
	r := gin.Default()

	// setup routes and middleware
	setupMiddleware(r)
	setupRoutes(r)

	// start app
	if err := r.Run(); err != nil {
		fmt.Println("fatal error while run")
		return
	}
}

func setupRoutes(r *gin.Engine) {
	r.GET("/", handlers.Index)
}

func setupMiddleware(r *gin.Engine) {
}
