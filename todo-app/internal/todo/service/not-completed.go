package todo

import "db-study/pkg/models"

func (s *TodoService) TodosNotCompleted() ([]models.Todo, error) {

	return s.repository.TodosNotCompleted()
}
