package psqlstorage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
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

func (a *Archiver) SaveBatch(messages []redis.XMessage) error {
	tx, err := a.db.dbpool.Begin(context.Background())
	if err != nil {
		return err
	}

	// https://github.com/jackc/pgx/blob/master/batch_test.go
	batch := &pgx.Batch{}

	for _, msg := range messages {
		newStruct, err := structpb.NewStruct(msg.Values)
		if err != nil {
			// release the tx
			return err
		}

		json, err := newStruct.MarshalJSON()
		if err != nil {
			// release the tx
			return err
		}

		batch.Queue("insert into train_archive(sequence_id, data) values($1, $2)", msg.ID, json)
	}

	br := tx.SendBatch(context.Background(), batch)

	for range messages {
		// call this for each item in the batch queue to get the results for them
		ct, err := br.Exec()
		if err != nil {
			return err
		}
		log.Println(ct.RowsAffected())
	}

	err = br.Close()
	if err != nil {
		// release the tx
		return err
	}

	return tx.Commit(context.Background())
}
