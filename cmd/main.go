package main

import (
	"api-go-gin/internal/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//init
	// Inicializar configurações
	//cfg := config.NewConfig()

	r := gin.Default()

	// Serve static files
	r.Static("/static", "./web/static")

	// Serve templates
	r.LoadHTMLGlob("web/templates/*")

	// Route for the form
	r.GET("/", controller.ShowForm)
	r.POST("/submit", controller.SubmitForm)

	log.Fatal(r.Run(":8080"))
}
