package main

import (
	"courseCRUD/cockroach"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func main() {
	err := Up()

	if err != nil {
		log.Panic(err)
	}
}

func Up() error {
	db, err := sql.Open("postgres", cockroach.GetDatabaseUrl())

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	err = goose.Up(db, "./migrations")

	return err
}

func Down() error {
	db, err := sql.Open("postgres", cockroach.GetDatabaseUrl())

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	err = goose.Down(db, "./migrations")

	return err
}
