package database

import (
	"errors"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	*sqlx.DB
	driver string
	url    string
}

func New() *Database {
	driver := os.Getenv("DATABASE_DRIVER")
	url := os.Getenv("DATABASE_URL")

	return &Database{
		driver: driver,
		url:    url,
	}
}

func (d *Database) Run() error {
	var err error
	d.DB, err = sqlx.Open(d.driver, d.url)
	if err != nil {
		return errors.New("Problem with Postgres " + err.Error())
	}

	err = d.Ping()
	if err != nil {
		return errors.New("Problem trying to Ping Postgres" + err.Error())
	}

	log.Println("Successful connection with PlayService Postgres")
	return nil
}
