package repository

import (
	"github.com/jmoiron/sqlx"
)

type TodoRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *TodoRepository {
	r := TodoRepository{
		db: db,
	}

	return &r
}
