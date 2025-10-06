package repository

import (
	"db-study/pkg/dto"
	"db-study/pkg/models"
	"fmt"
)

func (r *TodoRepository) AddTodo(data dto.CreateTodo) (models.Todo, error) {

	var createdTodo models.Todo

	row := r.db.QueryRowx(fmt.Sprintf("INSERT INTO todos (title, text) values ('%v', '%v') RETURNING id, title, text, time, completed", data.Title, data.Text))

	if row.Err() != nil {
		return models.Todo{}, row.Err()
	}

	err := row.StructScan(&createdTodo)

	if err != nil {
		return models.Todo{}, err
	}

	return createdTodo, nil
}

func (r *TodoRepository) MarkAsCompleted(todoId int) (models.Todo, error) {
	var createdTodo models.Todo

	row := r.db.QueryRowx(fmt.Sprintf("UPDATE todos SET completed=true WHERE id=%d RETURNING id, title, text, time, completed", todoId))

	err := row.StructScan(&createdTodo)

	if err != nil {
		return models.Todo{}, err
	}

	return createdTodo, nil
}
