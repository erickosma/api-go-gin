package http

import (
	"api-go-gin/internal/domain"
	"api-go-gin/internal/repository"
	"api-go-gin/internal/usecase"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(router *gin.Engine, db *mongo.Database) {
	repo := repository.NewMongoUserRepository(db)
	userUsecase := usecase.NewUserUsecase(repo)
	handler := &UserHandler{userUsecase}

	router.POST("/users", handler.CreateUser)
	router.GET("/users/:id", handler.GetUserByID)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.usecase.CreateUser(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.usecase.GetUserByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
