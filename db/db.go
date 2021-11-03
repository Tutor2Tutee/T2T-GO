package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const ConnectionTimeOut = 5

type Resource struct {
	DB *mongo.Database
}

// init will set prefix
func init() {
	log.SetPrefix("[MongoDB] ")
}

// GetResource will create Resource of given db server
func GetResource(UserId string, Password string, DbUrl string, DbName string) *Resource {
	var clientOptions *options.ClientOptions
	if UserId != "" {
		credential := options.Credential{
			Username: UserId,
			Password: Password,
		}
		clientOptions = options.Client().ApplyURI(DbUrl).SetAuth(credential)
	} else {
		clientOptions = options.Client().ApplyURI(DbUrl)
	}

	mClient, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalln(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), ConnectionTimeOut*time.Second)
	defer cancel()
	if err := mClient.Connect(ctx); err != nil {
		log.Fatalln(err.Error())
	}
	return &Resource{DB: mClient.Database(DbName)}
}
