package todo

import (
	"db-study/pkg/dto"
	"db-study/pkg/models"
	"db-study/pkg/validations"
)

func (s *TodoService) AddTodo(todo dto.CreateTodo) (models.Todo, error) {
	valErr := validations.CreateTodoValidation(todo.Title)

	if valErr != nil {
		return models.Todo{}, valErr
	}

	createdTodo, err := s.repository.AddTodo(todo)

	if err != nil {
		return models.Todo{}, err
	}

	return createdTodo, nil
}
