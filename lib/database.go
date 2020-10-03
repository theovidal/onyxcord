package lib

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Client

func OpenDatabase() {
	uri := fmt.Sprint("mongodb://", GlobalConfig.Database.Address, ":", GlobalConfig.Database.Port)
	var err error
	Db, err = mongo.NewClient(
		options.Client().ApplyURI(uri).SetAuth(options.Credential{
			Username:   GlobalConfig.Database.Username,
			Password:   GlobalConfig.Database.Password,
			AuthSource: GlobalConfig.Database.AuthSource,
		}),
	)
	if err != nil {
		panic(err)
	}

	err = Db.Connect(context.Background())
	if err != nil {
		panic(err)
	}
}
