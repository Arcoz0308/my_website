package database

import (
	"context"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client
var Paste *mongo.Collection

func Init() {
	clientOptions := options.Client().ApplyURI(utils.Config.Database.Url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	Client = client
	Paste = client.Database("website").Collection("arcpaste")
}
