package todo

import "db-study/pkg/models"

func (s *TodoService) TodosCompleted() ([]models.Todo, error) {

	return s.repository.TodosCompleted()
}
