package category

func (s *CategoryService) DeleteCategory(categoryId int) error {

	return s.repository.DeleteCategory(categoryId)

}
