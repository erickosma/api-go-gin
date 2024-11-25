package controller

import (
	"api-go-gin/internal/domain"
	"api-go-gin/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ShowForm(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

func SubmitForm(c *gin.Context) {
	time.Sleep(3 * time.Second)
	var data domain.User
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send WhatsApp message
	if err := service.SendWhatsAppMessage(data.WhatsApp, "Obrigado por se cadastrar, "+data.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send WhatsApp message"})
		return
	}

	// Save data to Google Sheets
	if err := service.SaveToGoogleSheets(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save to Google Sheets"})
		return
	}

	c.HTML(http.StatusOK, "done.html", nil)
}
