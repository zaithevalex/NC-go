package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbHost   = "localhost"
	dbName   = "db"
	dbPass   = "pass"
	dbUser   = "user"
)

const AddToDb = "insert into note(id, content) values(%d, '%s');"
const GetDbSize = "select count(*) from note;"
const GetFromDb = "select * from note where id = %d;"

type Content struct {
	Body string `json:"body"`
}

type Note struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

func Connection() (*sqlx.DB, error) {
	db, err := sqlx.Connect(dbDriver, fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s", dbUser, dbName, dbPass, dbHost))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Get(id int, db *sqlx.DB) (*Note, error) {
	var note Note
	err := db.Get(&note, fmt.Sprintf(GetFromDb, id))
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func Insert(content Content, db *sqlx.DB) (int, error) {
	var size int
	err := db.Get(&size, fmt.Sprintf(GetDbSize))
	if err != nil {
		return -1, err
	}

	_, err = db.Exec(fmt.Sprintf(AddToDb, size+1, content.Body))
	if err != nil {
		return -1, err
	}

	return size + 1, nil
}
