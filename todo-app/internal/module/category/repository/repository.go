package category

import "github.com/jmoiron/sqlx"

type CategoryRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *CategoryRepository {

	r := CategoryRepository{
		db: db,
	}

	return &r

}
