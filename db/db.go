package db

import (
	"context"
	"ethertemp/models"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var mongoClient *mongo.Client
var tempsCollection *mongo.Collection

func InitializeMongo() error {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatalln("error reading uri value for mongo")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	mongoClient = client
	return err
}

func AddTemperature(temperature float32) {
	if tempsCollection == nil {
		tempsCollection = mongoClient.Database("ethtemp").Collection("temperatures")
	}

	model := models.Tmp{
		Temperature: temperature,
		Timestamp:   time.Now(),
	}

	_, err := tempsCollection.InsertOne(context.TODO(), model)
	if err != nil {
		fmt.Println("error inserting a value:", err.Error())
	}
}

func Close() {
	_ = mongoClient.Disconnect(context.TODO())
}
