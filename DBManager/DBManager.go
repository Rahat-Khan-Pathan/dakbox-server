package DBManager

import (
	"context"
	"log"

	// get an object type
	// "encoding/json"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbURL string = "mongodb+srv://dakbox:XJ2WJt5Tv7OUN3QO@cluster0.2spjz.mongodb.net/Cluster0?retryWrites=true&w=majority"
var SystemCollections RTXCollections

type RTXCollections struct {
	SampleCollection *mongo.Collection
	Messages         *mongo.Collection
}

func InitRTXCollections() bool {
	var err error
	SystemCollections.SampleCollection, err = GetMongoDbCollection("rtx_db", "sample_collection")
	if err != nil {
		return false
	}

	SystemCollections.Messages, err = GetMongoDbCollection("dakbox", "messages")
	if err != nil {
		return false
	}

	return err == nil
}

// GetMongoDbConnection get connection of mongodb
func getMongoDbConnection() (*mongo.Client, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://dakbox:IEPxCNYjSFhzu5uN@cluster0.2spjz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client, err := getMongoDbConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)

	return collection, nil
}
