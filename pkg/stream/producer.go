package stream

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"log"
	"redis-trains/gen/pb/train_stream_pb"
	"time"
)

type Producer struct {
	host     string
	password string
	port     int
	db       int
	redis    *redis.Client
	stream   string
	maxLen   int64
	approx   bool

	closingChan chan bool
}

func NewProducer(host, password string, port, db int, stream string, maxLen int64, approx bool) (*Producer, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	producer := Producer{
		host:        host,
		password:    password,
		port:        port,
		db:          db,
		redis:       rdb,
		stream:      stream,
		maxLen:      maxLen,
		approx:      approx,
		closingChan: make(chan bool),
	}

	go producer.heartBeat()

	return &producer, nil
}

func (p *Producer) heartBeat() {
	for {
		select {
		case <-p.closingChan:
			return
		case <-time.After(1 * time.Second):
			err := p.redis.Ping(context.Background()).Err()
			if err != nil {
				newRedis := redis.NewClient(&redis.Options{
					Addr:     fmt.Sprintf("%s:%d", p.host, p.port),
					Password: p.password,
					DB:       p.db,
				})

				// todo do better
				err = newRedis.Ping(context.Background()).Err()
				if err != nil {
					log.Println(err)
					continue
				}

				p.redis = newRedis
			}
		}
	}
}

func (p *Producer) Produce(ctx context.Context, event *train_stream_pb.Event) (string, error) {
	bytes, err := proto.Marshal(event)
	if err != nil {
		return "", err
	}

	encoded := hex.EncodeToString(bytes)

	return p.redis.XAdd(ctx, &redis.XAddArgs{
		Stream: p.stream,
		MaxLen: p.maxLen,
		Approx: p.approx,
		Values: map[string]interface{}{"data": encoded},
	}).Result()
}

func (p *Producer) Close() error {
	p.closingChan <- true
	return p.redis.Close()
}
