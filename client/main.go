package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/arshabbir/sensormongogrpc/domain/sensorpb"
	"google.golang.org/grpc"
)

func main() {
	host := "localhost:9091"
	if err := startClient(host); err != nil {
		log.Println("Error Starting the client")
		return
	}
}

func startClient(host string) error {

	con, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		return err
	}

	c := sensorpb.NewSensorServiceClient(con)

	stream, err := c.SendData(context.Background())
	if err != nil {
		return err
	}

	for {

		if err := stream.Send(&sensorpb.SensorReq{Id: rand.Int31(),
			Reading: rand.Float32()}); err != nil {
			log.Println("Error sending data")
			break
		}

		time.Sleep(time.Second)
	}

	return nil
}
