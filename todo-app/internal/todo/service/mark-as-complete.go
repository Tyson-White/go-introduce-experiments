package todo

import "db-study/pkg/models"

func (s *TodoService) MarkAsCompleted(todoId int) (models.Todo, error) {

	todo, err := s.repository.MarkAsCompleted(todoId)

	if err != nil {
		return todo, err
	}

	return todo, nil

}
