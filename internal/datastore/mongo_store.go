package datastore

import (
	"context"
	"log"
	"time"

	"github.com/SimplQ/simplQ-golang/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// A Structure with Collections frequently used and a pointer to the client
type MongoDB struct {
    Client *mongo.Client
    Queue *mongo.Collection
    Token *mongo.Collection
}

var mongodb MongoDB

func ConnectMongoDB(uri string) *MongoDB {
    log.Println("Connection to MongoDB...")

    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
    
    if err != nil {
        log.Fatal(err)
    }

    mongodb = MongoDB {
        client,
        client.Database("simplq").Collection("queue"),
        client.Database("simplq").Collection("token"),
    }

    log.Println("Successfully connected to MongoDB!")

    return &mongodb
}

func (mongodb MongoDB) CreateQueue(queue models.Queue) models.QueueId {
    queue_insert := bson.D{
        {"queueName", queue.QueueName},
        {"isPaused", false},
        {"isDeleted", false},
        {"creationTime", time.Now()},
        {"deletionTime", time.Now()},
    }

    result, err := mongodb.Queue.InsertOne(context.TODO(), queue_insert)

    if err != nil {
        log.Fatal(err)
    }
 
    stringId := result.InsertedID.(primitive.ObjectID).Hex()

    return models.QueueId(stringId)
}
