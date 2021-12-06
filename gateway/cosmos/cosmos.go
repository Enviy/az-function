package cosmos

import (
	"os"
	"fmt"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson"
)

// Gateway defines mongoDB interaction methods
type Gateway interface {
	Update(databaseName, collectionName string, doc interface{}) (*mongo.InsertOneResult, error)
}

type gateway struct {
	client *mongo.Client
	ctx *context.Context
}

// New return new mongoDB Gateway
func New() (Gateway, error) {
	connectionString := os.Getenv("AZ_CONNECTIONSTRING")
	opts := options.Client().ApplyURI(connectionString)
	client, err := mongo.NewClient(opts)
	if err != nil {
		fmt.Printf("Unable to instantiate new mongoDB client: %v", err)
		return nil, err
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Unable to create client connection: %v", err)
		return nil, err
	}
	// TODO: Correct Ping() input args
	// err = client.Ping(ctx)
	return &gateway{
		client: client,
	}, nil
}

// Update updates mongoDB collection.
func (g *gateway) Update(databaseName, collectionName string, doc interface{}) (*mongo.InsertOneResult, error) {
	// database := g.client.Database(databaseName)
	collection := g.client.Database(databaseName).Collection(collectionName)
	result, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("Unable to write doc to mongoDB: %v", err)
		return nil, err
	}
	return result, nil
}
