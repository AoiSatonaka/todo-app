package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectClient() (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASS"),
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
		os.Getenv("MONGO_OPTIONS"),
	)
	fmt.Println(uri)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	cli, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	return cli, err
}

func closeClient(c *mongo.Client) error {
	if err := c.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}

//TODO:動作確認後削除
func Initc()(*mongo.Client,error){
  return connectClient()
}
func closec(c *mongo.Client)(error){
  return closeClient(c)
}
