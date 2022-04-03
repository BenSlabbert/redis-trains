package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"os/signal"
	"redis-trains/pkg/psqlstorage"
	"redis-trains/pkg/stream"
	"syscall"
)

const (
	Start        = "0-0"
	Stream       = "train-events"
	Group        = "archiver"
	Consumer     = "archiver-1"
	LastConsumed = ">"
	BatchSize    = 100
)

// GitCommit is set during compilation
var GitCommit string

func main() {
	fmt.Printf("GitCommit: %s", GitCommit)

	archiver, err := psqlstorage.NewArchiver()
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if e := archiver.Close(); e != nil {
			log.Println(e)
		}
	}()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer func() {
		if e := rdb.Close(); e != nil {
			log.Println(e)
		}
	}()

	var consumer *stream.Consumer
	consumer, err = stream.NewConsumer(context.Background(), rdb, Stream, Group, Start, LastConsumed, Consumer)
	if err != nil {
		log.Fatalln(err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-sigs:
			// os.Exit does not run defer funcs
			return
		default:
			// continue below
		}

		var messages []redis.XMessage
		messages, err = consumer.Consume(BatchSize)
		if err != nil && err != redis.Nil {
			log.Fatalln(err)
		}

		if len(messages) == 0 {
			continue
		}

		msgIds := make([]string, len(messages))
		for i := range msgIds {
			msgIds[i] = messages[i].ID
		}

		err = archiver.SaveBatch(messages)
		if err != nil {
			log.Fatalln(err)
		}

		err = rdb.XAck(context.Background(), Stream, Group, msgIds...).Err()
		if err != nil {
			log.Fatalln(err)
		}
	}
}
