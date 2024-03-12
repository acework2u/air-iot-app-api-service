package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	fmt.Println("<!------------DB URL--------------->\n", EnvMongoURI())
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))

	if err != nil {
		log.Fatal(err)

	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the Database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Connected to DB")

	return client

}
func ConnectDB2() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvDB2()))

	if err != nil {
		log.Fatal(err)

	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the Database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Connected to DB")

	return client

}

// getting datatbase collections

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	collection := client.Database("airs").Collection(collectionName)

	return collection
}
