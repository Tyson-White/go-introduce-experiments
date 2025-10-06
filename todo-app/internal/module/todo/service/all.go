package todo

import "db-study/pkg/models"

func (s *TodoService) Todos() ([]models.Todo, error) {
	todos, err := s.repository.Todos()

	if err != nil {
		return todos, err
	}

	return todos, nil
}
