package redisstorage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type TrainRoute struct {
	Origin      string `json:"origin,omitempty"`
	Destination string `json:"destination,omitempty"`
}

func (tr *TrainRoute) SwitchDirection() {
	tmp := tr.Origin
	tr.Origin = tr.Destination
	tr.Destination = tmp
}

type KVStore struct {
	host        string
	password    string
	port        int
	db          int
	redis       *redis.Client
	closingChan chan bool
}

func NewKVStore(host, password string, port, db int) (*KVStore, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	kv := &KVStore{
		host:        host,
		password:    password,
		port:        port,
		db:          db,
		redis:       rdb,
		closingChan: make(chan bool),
	}

	go kv.heartBeat()

	return kv, nil
}

func (kvs *KVStore) heartBeat() {
	for {
		select {
		case <-kvs.closingChan:
			return
		case <-time.After(1 * time.Second):
			err := kvs.redis.Ping(context.Background()).Err()
			if err != nil {
				newRedis := redis.NewClient(&redis.Options{
					Addr:     fmt.Sprintf("%s:%d", kvs.host, kvs.port),
					Password: kvs.password,
					DB:       kvs.db,
				})

				// todo do better
				err = newRedis.Ping(context.Background()).Err()
				if err != nil {
					log.Println(err)
					continue
				}

				kvs.redis = newRedis
			}
		}
	}
}

func (kvs *KVStore) Close() error {
	kvs.closingChan <- true
	return kvs.redis.Close()
}

func (kvs *KVStore) GetTrainRoute(train string) (*TrainRoute, error) {
	exists, err := kvs.redis.HExists(context.Background(), "train-route", train).Result()
	if err != nil {
		return nil, err
	}

	if !exists {
		tr := &TrainRoute{
			Origin:      "Kings Cross",
			Destination: "Kentish Town",
		}

		marshal, err := json.Marshal(tr)
		if err != nil {
			return nil, err
		}
		err = kvs.redis.HSet(context.Background(), "train-route", train, marshal).Err()
		if err != nil {
			return nil, err
		}
	}

	get, err := kvs.redis.HGet(context.Background(), "train-route", train).Bytes()
	if err != nil {
		return nil, err
	}

	tr := new(TrainRoute)
	err = json.Unmarshal(get, tr)
	if err != nil {
		return nil, err
	}

	return tr, nil
}
