package psqlstorage

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func init() {
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/train_archive?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// todo embed these files
	m, err := migrate.NewWithDatabaseInstance(
		"file:///home/ben/Goland/redis-trains/pkg/psqlstorage/migration/archive",
		"postgres", driver)
	if err != nil {
		log.Fatalln(err)
	}

	// number of files we need to run
	err = m.Up()

	if err != nil {
		if err.Error() == "no change" {
			log.Println("migration: no db changes to be made")
			return
		}
		log.Fatalln(err)
	}

	log.Println("migration: done")
}
