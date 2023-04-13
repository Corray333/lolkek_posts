package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var MG MongoInstance

func Connect() error {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("No .env file found!")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		fmt.Println("MONGODB_URI is empty!")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database("posts_db")

	if err != nil {
		return err
	}

	MG = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}
