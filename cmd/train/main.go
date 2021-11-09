package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
	"redis-trains/pkg/stream"
)

var ctx = context.Background()

const Start = "0-0"
const Stream = "train-events"
const Group = "group-1"
const Consumer = "consumer-1"
const LastConsumed = ">"
const BatchSize = 10

func main() {
	ExampleClient()
}

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	producer := stream.NewProducer(rdb, Stream, 1000, true)
	s, err := structpb.NewStruct(map[string]interface{}{})
	if err != nil {
		log.Fatalln(err)
	}

	s.Fields["test"] = structpb.NewStringValue("value")
	newId, err := producer.Produce(ctx, s)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(newId)

	consumer, err := stream.NewConsumer(ctx, rdb, Stream, Group, Start, LastConsumed, Consumer)
	if err != nil {
		log.Fatalln(err)
	}

	messages, err := consumer.Consume(BatchSize)
	if err != nil {
		log.Fatalln(err)
	}

	if len(messages) == 0 {
		messages, err = consumer.Consume(BatchSize)
		if err != nil {
			log.Fatalln(err)
		}
	}

	for _, msg := range messages {
		log.Println(msg.ID)
		cnt, err := rdb.XAck(ctx, Stream, Group, msg.ID).Result()
		if err != nil && err != redis.Nil {
			log.Fatalln(err)
		}
		log.Println(cnt)
	}
}
