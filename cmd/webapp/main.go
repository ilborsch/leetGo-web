package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	handlers "github.com/ilborsch/leetGo-web/internal/handlers/index"
)

func main() {
	// init config

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
