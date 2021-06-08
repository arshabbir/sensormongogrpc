package client

import (
	"context"
	"log"
	"time"

	"github.com/arshabbir/sensormongogrpc/domain/sensorpb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbClient struct {
	client *mongo.Client
}

type DBClient interface {
	Insert(*sensorpb.SensorReq) error
}

func NewDBClient() DBClient {

	//Initialize the mongodb connection object

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println("Error creating mongo client ", err)
		return nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println("Error connecting to DB", err)
		return nil
	}
	//defer client.Disconnect(ctx)

	return &dbClient{client: client}
}
func (d *dbClient) Insert(data *sensorpb.SensorReq) error {

	log.Printf("dbClient: Data : %v", data)

	c := d.client.Database("sensordb").Collection("sensordata")

	result, err := c.InsertOne(context.Background(), data)
	if err != nil {
		log.Println("Insertion Error", err)
		return err
	}

	log.Println("Successfully Inserted. Record : ", result.InsertedID)

	return nil
}
