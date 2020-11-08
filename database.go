package onyxcord

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
		panic(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	return
}
