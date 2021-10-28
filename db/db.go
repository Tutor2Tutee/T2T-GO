package db

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
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
func GetResource() *Resource {
	UserId := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASS")
	DbUrl := os.Getenv("DB_URL")
	DbName := os.Getenv("DB_NAME")
	credential := options.Credential{
		Username: UserId,
		Password: Password,
	}
	clientOptions := options.Client().ApplyURI(DbUrl).SetAuth(credential)

	mClient, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalln(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), ConnectionTimeOut*time.Second)
	defer cancel()
	if err := mClient.Connect(ctx); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("MongoDB connected")
	return &Resource{DB: mClient.Database(DbName)}
}
