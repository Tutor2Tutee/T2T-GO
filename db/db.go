package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// init will set prefix
func init() {
	log.SetPrefix("[MongoDB] ")
}

// MongoConn will create a connection to mongo database with given arguments
func MongoConn(UserId string, Password string, DbUrl string) *mongo.Client {
	credential := options.Credential{
		Username: UserId,
		Password: Password,
	}
	clientOptions := options.Client().ApplyURI(DbUrl).SetAuth(credential)

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
