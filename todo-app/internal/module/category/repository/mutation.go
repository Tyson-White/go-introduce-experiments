package category

import (
	"db-study/pkg/models"
	"log"
)

func (r *CategoryRepository) CreateCategory(name string) (models.Category, error) {
	var category models.Category

	row := r.db.QueryRowx("INSERT INTO category (name) values ($1) RETURNING id, name", name)

	err := row.StructScan(&category)

	if err != nil {
		log.Println("[ERROR] Repository - category:", err)
		return category, err
	}

	return category, nil
}
