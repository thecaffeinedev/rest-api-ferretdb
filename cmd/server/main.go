package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/thecaffeinedev/rest-api-ferretdb/internal/configs"
	"github.com/thecaffeinedev/rest-api-ferretdb/internal/database"
	transportHTTP "github.com/thecaffeinedev/rest-api-ferretdb/internal/transport/http"
	"github.com/thecaffeinedev/rest-api-ferretdb/internal/users"
)

// Run - sets up our app
func Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Setting Up Our APP")

	cfg := configs.GetConfig()

	db := database.NewDatabase(cfg)

	userService := users.NewService(db)

	handler := transportHTTP.NewHandler(userService)

	if err := handler.Serve(); err != nil {
		log.Error("failed to gracefully serve our application")
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Error(err)
		log.Fatal("Error starting up the Server")
	}
}
