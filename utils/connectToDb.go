package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectToDb() func() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))
	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Conneted to Db")
	return func() {
		cancel()
		if err = Client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}
}
