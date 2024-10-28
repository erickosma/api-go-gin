package main

import (
	"api-go-gin/config"
	"api-go-gin/internal/handler"
	"api-go-gin/pkg/mongo"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	db, err := mongo.ConnectMongo(config.LoadConfig())
	if err != nil {
		log.Fatal(err)
	}

	http.NewUserHandler(r, db)

	if err := r.Run(); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
