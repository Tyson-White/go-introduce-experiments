package todo

import "db-study/pkg/models"

func (s *TodoService) MarkAsUncompleted(todoId int) (models.Todo, error) {

	todo, err := s.repository.MarkAsUncompleted(todoId)

	if err != nil {
		return todo, err
	}

	return todo, nil

}
