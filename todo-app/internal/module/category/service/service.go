package category

import category "db-study/internal/module/category/repository"

type CategoryService struct {
	repository *category.CategoryRepository
}

func New(repo *category.CategoryRepository) *CategoryService {

	s := CategoryService{
		repository: repo,
	}

	return &s

}
