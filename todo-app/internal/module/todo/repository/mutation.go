package repository

import (
	"db-study/pkg"
	"db-study/pkg/dto"
	"db-study/pkg/models"
)

func (r *TodoRepository) AddTodo(data dto.CreateTodo) (models.Todo, error) {

	var createdTodo models.Todo

	row := r.db.QueryRowx("INSERT INTO todos (title, text, category_id) values ($1, $2, $3) RETURNING id, title, text, time, completed, category_id", data.Title, data.Text, data.Category)

	err := row.StructScan(&createdTodo)

	if err != nil {
		return models.Todo{}, err
	}

	return createdTodo, nil
}

func (r *TodoRepository) MarkAsCompleted(todoId int) (models.Todo, error) {
	var createdTodo models.Todo

	row := r.db.QueryRowx("UPDATE todos SET completed=true WHERE id=$1 RETURNING id, title, text, time, completed", todoId)

	err := row.StructScan(&createdTodo)

	if err != nil {
		return models.Todo{}, err
	}

	return createdTodo, nil
}

func (r *TodoRepository) DeleteTodo(todoId int) error {

	res, err := r.db.Exec("DELETE FROM todos WHERE id=$1", todoId)

	if err != nil {
		return err
	}

	n, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if n == 0 {
		return pkg.ErrNotFound
	}

	return nil
}
