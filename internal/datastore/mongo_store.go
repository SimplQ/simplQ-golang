package datastore

import (
	"context"
	"log"
    "os"

	"github.com/SimplQ/simplQ-golang/internal/models"

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

func NewMongoDB() *MongoDB {
    // Use local mongodb instance if env variable not set
    uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
    
    if val, ok := os.LookupEnv("MONGO_URI"); ok {
        uri = val
	}

    log.Println("Connecting to MongoDB...")

    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
    
    if err != nil {
        log.Fatal(err)
    }

    new_mongodb := MongoDB {
        client,
        client.Database("simplq").Collection("queue"),
        client.Database("simplq").Collection("token"),
    }

    log.Println("Successfully connected to MongoDB!")

    return &new_mongodb
}

func (mongodb MongoDB) CreateQueue(queue models.Queue) models.QueueId {
    // Set id to empty so its generated by mongoDB
    queue.Id = ""

    result, err := mongodb.Queue.InsertOne(context.TODO(), queue)

    if err != nil {
        log.Fatal(err)
    }
 
    stringId := result.InsertedID.(primitive.ObjectID).Hex()

    return models.QueueId(stringId)
}

func (mongodb MongoDB) ReadQueue(models.QueueId) models.Queue {
    panic("Not implemented")
}

func (mongodb MongoDB) PauseQueue(models.QueueId) {
    panic("Not implemented")
}

func (mongodb MongoDB) ResumeQueue(models.QueueId) {
    panic("Not implemented")
}

func (mongodb MongoDB) DeleteQueue(models.QueueId) {
    panic("Not implemented")
}

func (mongodb MongoDB) AddTokenToQueue(models.QueueId, models.Token) {
    panic("Not implemented")
}

func (mongodb MongoDB) ReadToken(models.TokenId) {
    panic("Not implemented")
}

func (mongodb MongoDB) RemoveToken(models.TokenId) {
    panic("Not implemented")
}
