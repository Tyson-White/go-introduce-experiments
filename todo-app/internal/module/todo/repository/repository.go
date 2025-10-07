package repository

import (
	"github.com/jmoiron/sqlx"
)

type TodoRepository struct {
	db *sqlx.DB
}

const schema = `
	CREATE TABLE IF NOT EXISTS todos (
		id serial not null unique, 
		title varchar(255) not null, 
		text text, 
		time timestamp default now(), 
		completed bool default false,
		category_id int REFERENCES category(id)
	)`

func New(db *sqlx.DB) *TodoRepository {
	r := TodoRepository{
		db: db,
	}

	_, err := db.Exec(schema)

	if err != nil {
		panic(err)
	}

	return &r
}
