package main

import (
	"errors"
	"log"
	"net"

	"github.com/arshabbir/sensormongogrpc/domain/sensorpb"
	sensorservice "github.com/arshabbir/sensormongogrpc/server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	host := "0.0.0.0:9091"

	if err := startServerApp(host); err != nil {
		log.Println("Error starting server  : ", err)
		return
	}

	return
}

func startServerApp(host string) error {

	lis, err := net.Listen("tcp", host)

	if err != nil {
		return err
	}

	service := sensorservice.NewSensorService()

	if service == nil {
		return errors.New("Error")
	}
	s := grpc.NewServer()

	sensorpb.RegisterSensorServiceServer(s, service)

	log.Println("Starting the gRPC server")

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		return err
	}

	log.Println("gRPC Server Started")

	return nil
}
