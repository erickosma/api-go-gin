package usecase

import (
	"api-go-gin/internal/domain"
	"api-go-gin/internal/repository"
	"context"
)

type UserUsecase struct {
	repo *repository.MongoUserRepository
}

func NewUserUsecase(repo *repository.MongoUserRepository) *UserUsecase {
	return &UserUsecase{repo}
}

func (u *UserUsecase) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	return u.repo.CreateUser(ctx, user)
}

func (u *UserUsecase) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	return u.repo.GetUserByID(ctx, id)
}
