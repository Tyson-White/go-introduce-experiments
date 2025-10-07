package category

import "github.com/jmoiron/sqlx"

type CategoryRepository struct {
	db *sqlx.DB
}

const schema = `
	CREATE TABLE IF NOT EXISTS category (
		id serial not null unique,
		name varchar(100) unique not null
	);
`

func New(db *sqlx.DB) *CategoryRepository {

	r := CategoryRepository{
		db: db,
	}

	db.Exec(schema)

	return &r

}
