package category

import (
	"db-study/pkg/models"
	"log"
)

func (r *CategoryRepository) CategoriesAll() ([]models.Category, error) {
	categories := []models.Category{}

	rows, err := r.db.Queryx("SELECT * FROM category")

	if err != nil {
		return categories, err
	}

	defer rows.Close()

	for rows.Next() {
		cat := models.Category{}

		err = rows.StructScan(&cat)

		if err != nil {
			log.Println("[ERROR] category reading: ", err)
			continue
		}

		categories = append(categories, cat)
	}

	return categories, nil
}
