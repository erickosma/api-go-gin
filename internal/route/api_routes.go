package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPIRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/", status)
	}
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
