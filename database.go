package onyxcord

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// OpenDatabase returns the mongodb database client (if needed and taking credentials from the configuration file)
func OpenDatabase(config *Config) (client *mongo.Client) {
	uri := fmt.Sprint("mongodb://", config.Database.Address, ":", config.Database.Port)
	var err error
	client, err = mongo.NewClient(
		options.Client().ApplyURI(uri).SetAuth(options.Credential{
			Username:   config.Database.Username,
			Password:   config.Database.Password,
			AuthSource: config.Database.AuthSource,
		}),
	)
	if err != nil {
		log.Panicf("‼ Error creating database entity: %s", err.Error())
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Panicf("‼ Error connecting to database: %s", err.Error())
	}

	return
}
