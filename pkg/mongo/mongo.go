package mongo

import (
	"api-go-gin/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const TimeOut = 10 // Defina o tempo limite desejado

func ConnectMongo(cfg config.Config) (*mongo.Database, error) {
	// Definir o contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), TimeOut*time.Second)
	defer cancel()

	// Criar e conectar o cliente MongoDB em uma única etapa
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Verifica a conexão com o MongoDB
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Retorna a conexão com o banco de dados especificado
	return client.Database(cfg.MongoDatabase), nil
}
