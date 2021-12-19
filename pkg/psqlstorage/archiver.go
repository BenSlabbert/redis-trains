package psqlstorage

import (
	"context"
	"encoding/hex"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4"
	"log"
	"strings"
)

type Archiver struct {
	db *DB
}

func NewArchiver() (*Archiver, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}

	return &Archiver{
		db: db,
	}, nil
}

func (a *Archiver) Close() error {
	a.db.dbpool.Close()
	return nil
}

func (a *Archiver) SaveBatch(messages []redis.XMessage) (err error) {
	var tx pgx.Tx
	tx, err = a.db.dbpool.Begin(context.Background())
	if err != nil {
		return err
	}

	// https://github.com/jackc/pgx/blob/master/batch_test.go
	batch := &pgx.Batch{}

	for _, msg := range messages {
		decodedBytes, err := hex.DecodeString(msg.Values["data"].(string))
		if err != nil {
			// release the tx
			return err
		}

		split := strings.Split(msg.ID, "-")

		batch.Queue("insert into train_archive(sequence_timestamp, sequence_increment, data) values($1, $2, $3)", split[0], split[1], decodedBytes)
	}

	br := tx.SendBatch(context.Background(), batch)

	for range messages {
		// call this for each item in the batch queue to get the results for them
		_, err = br.Exec()
		if err != nil {
			return err
		}
	}

	err = br.Close()
	if err != nil {
		// release the tx
		return err
	}

	log.Printf("saved: %d rows in batch", len(messages))
	return tx.Commit(context.Background())
}
