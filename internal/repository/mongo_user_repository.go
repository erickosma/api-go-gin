package repository

import (
	"api-go-gin/internal/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	db *mongo.Database
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{db}
}

func (r *MongoUserRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	collection := r.db.Collection("users")
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *MongoUserRepository) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	collection := r.db.Collection("users")
	objID, _ := primitive.ObjectIDFromHex(id)
	var user domain.User
	err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	return user, err
}

// Implementar os métodos de Update e Delete...
