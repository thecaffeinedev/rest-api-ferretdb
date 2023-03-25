package configs

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

// DBConfig describes the database config
type DBConfig struct {
	Collection string
	Name       string
	URI        string
}

// Config struct describes all the configs
type Config struct {
	Database *DBConfig
}

// GetConfig returns the config from environment variables
func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: error loading .env file, using environment values")
	}

	return &Config{
		Database: &DBConfig{
			Collection: os.Getenv("DatabaseCollection"),
			Name:       os.Getenv("DatabaseName"),
			URI:        os.Getenv("MONGO_URL"),
		},
	}
}
