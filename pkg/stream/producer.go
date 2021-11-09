package stream

import (
	"context"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/types/known/structpb"
)

type Producer struct {
	redis  *redis.Client
	stream string
	maxLen int64
	approx bool
}

func NewProducer(rdb *redis.Client, stream string, maxLen int64, approx bool) *Producer {
	return &Producer{
		redis:  rdb,
		stream: stream,
		maxLen: maxLen,
		approx: approx,
	}
}

func (p *Producer) Produce(ctx context.Context, values *structpb.Struct) (string, error) {
	return p.redis.XAdd(ctx, &redis.XAddArgs{
		Stream: p.stream,
		MaxLen: p.maxLen,
		Approx: p.approx,
		Values: values.AsMap(),
	}).Result()
}
