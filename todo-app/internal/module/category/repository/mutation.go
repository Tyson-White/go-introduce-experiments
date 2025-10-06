package category

import (
	"db-study/pkg/models"
	"fmt"
	"log"
)

func (r *CategoryRepository) CreateCategory(name string) (models.Category, error) {
	var category models.Category

	row := r.db.QueryRowx(fmt.Sprintf("INSERT INTO category (name) values ('%v') RETURNING id, name", name))

	err := row.StructScan(&category)

	if err != nil {
		log.Println("[ERROR] Repository - category:", err)
		return category, err
	}

	return category, nil
}
