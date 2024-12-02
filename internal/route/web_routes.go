package route

import (
	"api-go-gin/internal/controller"
	"github.com/gin-gonic/gin"
)

func RegisterWebRoutes(router *gin.Engine) {
	router.GET("/", controller.Index)
	router.GET("/form", controller.ShowForm)
	router.POST("/submit", controller.SubmitForm)
}
