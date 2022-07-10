package utils

import "os"

type Config struct {
	Environment string

	MongoUsername string
	MongoPassword string
	MongoURI      string
	MongoDBName   string
}

func NewConfig() *Config {
	config := &Config{
		Environment:   "dev",
		MongoUsername: "",
		MongoPassword: "",
		MongoURI:      "",
	}

	val := os.Getenv("ENV")
	if val != "" {
		config.Environment = val
	}

	val = os.Getenv("MONGO_USERNAME")
	if val != "" {
		config.MongoUsername = val
	}

	val = os.Getenv("MONGO_PASSWORD")
	if val != "" {
		config.MongoPassword = val
	}

	val = os.Getenv("MONGO_URI")
	if val != "" {
		config.MongoURI = val
	}

	val = os.Getenv("MONGO_DB_NAME")
	if val != "" {
		config.MongoDBName = val
	}

	return config
}
