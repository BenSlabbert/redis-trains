package main

import (
	"log"
	"os"
	"os/signal"
	"redis-trains/pkg/graph"
	"redis-trains/pkg/redisstorage"
	"redis-trains/pkg/stream"
	"redis-trains/pkg/train"
	"syscall"
)

const Stream = "train-events"

func main() {

	kvStore, err := redisstorage.NewKVStore("localhost", "", 6379, 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := kvStore.Close(); err != nil {
			log.Println(err)
		}
	}()

	rnc, err := graph.NewRailNetworkClient("bolt://localhost:7687", "neo4j", "neo4j", "", 2)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := rnc.Close(); err != nil {
			log.Println(err)
		}
	}()

	producer, err := stream.NewProducer("localhost", "", 6379, 0, Stream, 1000, true)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Println(err)
		}
	}()

	simple := train.NewSimple("Thomas", rnc, kvStore, producer)
	go simple.Run()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-sigs:
			// os.Exit does not run defer funcs
			log.Println("stopping train")
			simple.Stop()
			<-simple.Exiting
			log.Println("train stopped")
			return
		case err = <-simple.Exiting:
			if err == train.ErrTrainCompleted {
				// normal execution
				return
			}

			log.Printf("train ran into an issue: %v", err)
			return
		}
	}
}
