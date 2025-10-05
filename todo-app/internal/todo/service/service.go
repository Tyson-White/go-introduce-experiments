package todo

import (
	todo "db-study/internal/todo/repository"
)

type TodoService struct {
	repository *todo.TodoRepository
}

func New(repo *todo.TodoRepository) *TodoService {
	s := TodoService{
		repository: repo,
	}

	return &s
}
