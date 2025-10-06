package category

import "db-study/pkg/models"

func (s *CategoryService) CategoriesAll() ([]models.Category, error) {

	return s.repository.CategoriesAll()
}
