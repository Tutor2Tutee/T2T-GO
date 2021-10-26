package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// init will set prefix
func init() {
	log.SetPrefix("[MongoDB] ")
}

// MongoConn will create a connection to mongo database with os environment variable
func MongoConn() *mongo.Client {
	credential := options.Credential{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	}
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL")).SetAuth(credential)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
	}
	if pingErr := client.Ping(context.TODO(), nil); pingErr != nil {
		log.Fatalln(pingErr)
	}
	log.Println("MongoDB connected")
	return client
}
