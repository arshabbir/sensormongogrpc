package sensorservice

import (
	"log"

	"github.com/arshabbir/sensormongogrpc/domain/sensorpb"
	"github.com/arshabbir/sensormongogrpc/server/client"
)

type sservice struct {
	dbclient client.DBClient
}

type SService interface {
	Insert(*sensorpb.SensorReq) error
	SendData(stream sensorpb.SensorService_SendDataServer) error
}

func NewSensorService() SService {

	client := client.NewDBClient()

	if client == nil {
		return nil
	}
	return &sservice{dbclient: client}
}

func (s *sservice) SendData(stream sensorpb.SensorService_SendDataServer) error {

	for {
		data, err := stream.Recv()
		if err != nil {
			log.Println("Error receiving  data")
			return err
		}
		s.Insert(data)
	}
	return nil
}
func (s *sservice) Insert(data *sensorpb.SensorReq) error {
	//log.Println(data)

	return s.dbclient.Insert(data)
}
