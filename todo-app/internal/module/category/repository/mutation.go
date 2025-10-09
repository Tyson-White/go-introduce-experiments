package category

import (
	"db-study/pkg"
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

func (r *CategoryRepository) DeleteCategory(categoryId int) error {

	res, err := r.db.Exec("DELETE FROM category WHERE id = $1", categoryId)

	if err != nil {
		return err
	}

	n, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if n == 0 {
		return pkg.ErrNotFound
	}

	return nil

}
