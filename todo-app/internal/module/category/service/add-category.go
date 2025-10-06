package category

import (
	"db-study/pkg/models"
	"db-study/pkg/validations"
)

func (s *CategoryService) AddCategory(name string) (models.Category, error) {

	err := validations.CategoryCreateValidation(name)

	if err != nil {
		return models.Category{}, err
	}

	cat, err := s.repository.CreateCategory(name)

	if err != nil {
		return cat, err
	}

	return cat, nil

}
