package psqlstorage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
)

type DB struct {
	dbpool *pgxpool.Pool
}

func NewDB() (*DB, error) {
	dbpool, err := pgxpool.Connect(context.Background(), "postgresql://user:password@localhost:5432/train_archive?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &DB{
		dbpool: dbpool,
	}, nil
}

func (db *DB) ExecuteFile(path string) (err error) {
	var tx pgx.Tx
	tx, err = db.dbpool.Begin(context.Background())
	if err != nil {
		return err
	}

	var file []byte
	file, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), string(file))
	if err != nil {
		rollBackErr := tx.Rollback(context.Background())
		if rollBackErr != nil {
			return fmt.Errorf(err.Error() + " rollback err: " + rollBackErr.Error())
		}
		return err
	}

	return tx.Commit(context.Background())
}
