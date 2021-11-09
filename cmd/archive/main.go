package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"os/signal"
	"redis-trains/pkg/psqlstorage"
	"redis-trains/pkg/stream"
	"syscall"
)

const Start = "0-0"
const Stream = "train-events"
const Group = "archiver"
const Consumer = "archiver-1"
const LastConsumed = ">"
const BatchSize = 100

func main() {
	archiver, err := psqlstorage.NewArchiver()
	if err != nil {
		log.Fatalln(err)
	}

	err = archiver.SaveBatch([]redis.XMessage{
		{
			ID: "someId",
			Values: map[string]interface{}{
				"key": "value",
			},
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	err = archiver.Close()
	if err != nil {
		log.Fatalln(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	consumer, err := stream.NewConsumer(context.Background(), rdb, Stream, Group, Start, LastConsumed, Consumer)
	if err != nil {
		log.Fatalln(err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {

		select {
		case <-sigs:
			err = rdb.Close()
			if err != nil {
				log.Fatalln(err)
			}
			os.Exit(0)
		default:
			// continue below
		}

		var messages []redis.XMessage
		messages, err = consumer.Consume(BatchSize)
		if err != nil && err != redis.Nil {
			log.Fatalln(err)
		}

		log.Println(messages)

		// save as a batch
		//for _, msg := range messages {
		//}
	}
}
