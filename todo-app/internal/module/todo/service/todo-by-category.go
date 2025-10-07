package todo

import "db-study/pkg/models"

func (s *TodoService) TodoByCategory(category int) ([]models.Todo, error) {

	return s.repository.TodoByCategory(category)

}
