package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
	"os"
	"os/signal"
	"redis-trains/pkg/stream"
	"syscall"
	"time"
)

const Stream = "train-events"

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer func() {
		if err := rdb.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	producer := stream.NewProducer(rdb, Stream, 1000, true)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-sigs:
			os.Exit(0)
		case <-time.After(1 * time.Second):
			err := publishNewMessage(producer)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func publishNewMessage(producer *stream.Producer) error {
	s, err := structpb.NewStruct(map[string]interface{}{})
	if err != nil {
		return err
	}
	s.Fields["key"] = structpb.NewStringValue("value: " + time.Now().String())

	var newId string
	newId, err = producer.Produce(context.Background(), s)
	if err != nil {
		return err
	}
	log.Println(newId)
	return nil
}
