package database

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/thecaffeinedev/rest-api-ferretdb/internal/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database struct describes the database
type Database struct {
	Client     *mongo.Client
	Collection *mongo.Collection
	DBConfig   *configs.DBConfig
}

// NewDatabase returns a new database
func NewDatabase(cfg *configs.Config) *Database {
	client := connectDB(cfg.Database.URI)
	collection := getCollection(client, cfg.Database.Name, cfg.Database.Collection)
	return &Database{
		Client:     client,
		Collection: collection,
		DBConfig:   cfg.Database,
	}
}

// connectDB connects to the database and sets in config
func connectDB(dbURI string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	// d.Client = client
	return client
}

// getCollection gets database collections and sets in config
func getCollection(client *mongo.Client, dbName, dbCollection string) *mongo.Collection {
	return client.Database(dbName).Collection(dbCollection)
}
