package main

import (
	"api-go-gin/config"
	"api-go-gin/internal/route"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//init
	// Inicializar configurações
	cfg := config.NewConfig()

	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")
	// Serve templates
	r.LoadHTMLGlob("templates/*")
	// Initialize routes
	route.RegisterWebRoutes(r)
	route.RegisterAPIRoutes(r)

	log.Fatal(r.Run(":" + cfg.ServerPort))
}
