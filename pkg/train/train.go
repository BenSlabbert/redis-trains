package train

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"redis-trains/gen/pb/train_stream_pb"
	"redis-trains/pkg/redisstorage"
	"redis-trains/pkg/stream"
	"time"
)

var ErrTrainCompleted = errors.New("train has completed its route")
var ErrTrainStopped = errors.New("train has been stopped")

type State string

const Err State = "Err"
const Starting State = "Starting"
const ArrivingAtStation State = "ArrivingAtStation"
const LeavingStation State = "LeavingStation"
const TravelingToNextStation State = "TravelingToNextStation"
const Stopping State = "Stopping"
const TurningAround State = "TurningAround"
const Stopped State = "Stopped"

type Simple struct {
	Name    string
	Exiting chan error

	rnc      RailNetworkClient
	kvStore  *redisstorage.KVStore
	producer *stream.Producer
	state    State
	stop     bool
}

type RailNetworkClient interface {
	FindPath(origin, destination string) ([]string, error)
}

func NewSimple(name string, rnc RailNetworkClient, kvStore *redisstorage.KVStore, producer *stream.Producer) *Simple {
	return &Simple{Name: name, Exiting: make(chan error), rnc: rnc, kvStore: kvStore, producer: producer, state: Stopped}
}

func (s *Simple) Stop() {
	s.stop = true
}

func (s *Simple) Run() {
	route, err := s.kvStore.GetTrainRoute(s.Name)
	if err != nil {
		s.moveToErrState(err)
		return
	}

	cnt := 0

	for {
		if cnt == 2 {
			_ = s.moveTo(Stopping)
			_ = s.moveTo(Stopped)
			s.Exiting <- ErrTrainCompleted
			return
		}

		_ = s.moveTo(Starting)
		_ = s.runRoute(route)
		if err != nil {
			s.moveToErrState(err)
			return
		}

		_ = s.moveTo(TurningAround)
		route.SwitchDirection()
		cnt++
	}
}

func (s *Simple) runRoute(route *redisstorage.TrainRoute) error {
	path, err := s.rnc.FindPath(route.Origin, route.Destination)
	if err != nil {
		return err
	}

	for idx, station := range path {
		if idx == len(path)-1 {
			// we are at the end of the line
			err = s.moveTo(ArrivingAtStation)
			if err != nil {
				return err
			}
			log.Printf("arriving at station %s, this is the end of the line", station)
			return nil
		}

		err = s.moveTo(ArrivingAtStation)
		if err != nil {
			return err
		}
		log.Printf("arriving at station %s", station)

		err = s.moveTo(LeavingStation)
		if err != nil {
			return err
		}
		log.Printf("leaving station %s", station)

		err = s.moveTo(TravelingToNextStation)
		if err != nil {
			return err
		}
		nextStation := path[idx+1]
		log.Printf("travelling to station %s", nextStation)
	}

	return nil
}

func (s *Simple) moveToErrState(err error) {
	s.state = Err
	s.Exiting <- err
}

func (s *Simple) moveTo(state State) error {
	if s.stop {
		s.Exiting <- ErrTrainStopped
		return ErrTrainStopped
	}

	s.state = state

	sMap, err := structpb.NewStruct(map[string]interface{}{})
	if err != nil {
		return err
	}
	sMap.Fields["time"] = structpb.NewStringValue(time.Now().String())
	sMap.Fields["state"] = structpb.NewStringValue(string(state))

	e := &train_stream_pb.Event{
		Timestamp: timestamppb.Now(),
		Payload: &train_stream_pb.Event_Error{
			Error: &train_stream_pb.ErrorMessage{},
		},
	}

	nextId, err := s.producer.Produce(context.Background(), e)
	if err != nil {
		return err
	}

	log.Printf("produced new record: %s", nextId)

	<-time.After(1 * time.Second)
	return nil
}
