package stream

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type Consumer struct {
	redis           *redis.Client
	stream          string
	group           string
	consumer        string
	consumeFrom     string
	checkForPending bool
}

func NewConsumer(ctx context.Context, rdb *redis.Client, stream, group, start, consumeFrom, consumer string) (*Consumer, error) {
	err := rdb.XGroupCreateMkStream(ctx, stream, group, start).Err()
	if err != nil {
		// check if consumer group already exists for stream
		if err.Error() != "BUSYGROUP Consumer Group name already exists" {
			return nil, err
		}
	}

	return &Consumer{
		redis:           rdb,
		stream:          stream,
		group:           group,
		consumer:        consumer,
		consumeFrom:     consumeFrom,
		checkForPending: true,
	}, nil
}

func (c *Consumer) Consume(batchMaxSize int64) ([]redis.XMessage, error) {
	// look for pending messages
	consumeFrom := c.consumeFrom

	if c.checkForPending {
		pending, err := c.redis.XPending(context.Background(), c.stream, c.group).Result()
		if err != nil {
			return nil, err
		}

		if pending.Count > 0 && pending.Lower != "" {
			consumeFrom = pending.Lower
		}
		c.checkForPending = false
	}

	stream, err := c.redis.XReadGroup(context.Background(), &redis.XReadGroupArgs{
		Group:    c.group,
		Consumer: c.consumer,
		Streams:  []string{c.stream, consumeFrom},
		Count:    batchMaxSize,
		Block:    100 * time.Millisecond,
		NoAck:    false,
	}).Result()

	if err != nil {
		return nil, err
	}

	for _, xStream := range stream {
		if xStream.Stream == c.stream {
			return xStream.Messages, nil
		}
	}

	// should not happen
	return nil, errors.New("did not get messages for stream: " + c.stream)
}
