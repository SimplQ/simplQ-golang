package datastore

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
    Client *mongo.Client
    Queue *mongo.Collection
    Token *mongo.Collection
}

func ConnectMongoDB(uri string) MongoDB {
    log.Println("Connection to MongoDB...")

    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
    
    if err != nil {
        log.Fatal(err)
    }

    mongodb := MongoDB {
        client,
        client.Database("simplq").Collection("queue"),
        client.Database("simplq").Collection("token"),
    }

    log.Println("Successfully connected to MongoDB!")

    return mongodb
}
