package config

type Config struct {
	MongoURI      string
	MongoDatabase string
}

func LoadConfig() Config {
	return Config{
		MongoURI:      "mongodb://localhost:27017",
		MongoDatabase: "api_go",
	}
}
