package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	Products *mongo.Collection
)

func Init(url string, database string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)

	client_, err := mongo.Connect(context.Background(), opts) 
	if err != nil {
		return err
	}
	client = client_

	Products = client.Database(database).Collection("products")

	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err()
	return err
}

func Close() error {
	return client.Disconnect(context.Background())
}
