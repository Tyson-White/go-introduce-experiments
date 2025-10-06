package todo

import (
	"database/sql"
	"db-study/pkg/dto"
	"db-study/pkg/models"
	"fmt"
	"log"
)

type TodoRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *TodoRepository {
	r := TodoRepository{
		db: db,
	}

	return &r
}

func (r *TodoRepository) AddTodo(data dto.CreateTodo) (models.Todo, error) {

	var createdTodo models.Todo

	row := r.db.QueryRow(fmt.Sprintf("INSERT INTO todos (title, text) values ('%v', '%v') RETURNING id, title, text, time, completed", data.Title, data.Text))

	if row.Err() != nil {
		return models.Todo{}, row.Err()
	}

	err := row.Scan(&createdTodo.Id, &createdTodo.Title, &createdTodo.Text, &createdTodo.Time, &createdTodo.Completed)

	if err != nil {
		return models.Todo{}, err
	}

	return createdTodo, nil
}

func (r *TodoRepository) Todos() ([]models.Todo, error) {

	todos := []models.Todo{}

	rows, err := r.db.Query("SELECT * FROM todos")

	if err != nil {
		return todos, err
	}

	defer rows.Close()

	for rows.Next() {

		todo := models.Todo{}

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Text, &todo.Time, &todo.Completed)

		if err != nil {
			log.Println(err)
			continue
		}

		todos = append(todos, todo)
	}

	return todos, nil

}

func (r *TodoRepository) MarkAsCompleted(todoId int) (models.Todo, error) {
	var createdTodo models.Todo

	row := r.db.QueryRow(fmt.Sprintf("UPDATE todos SET completed=true WHERE id=%d RETURNING id, title, text, time, completed", todoId))

	err := row.Scan(&createdTodo.Id, &createdTodo.Title, &createdTodo.Text, &createdTodo.Time, &createdTodo.Completed)

	if err != nil {
		return models.Todo{}, err
	}

	return createdTodo, nil
}

func (r *TodoRepository) TodosNotCompleted() ([]models.Todo, error) {
	var todos = []models.Todo{}

	rows, err := r.db.Query("SELECT * FROM todos WHERE completed=false")

	if err != nil {
		return todos, err
	}

	defer rows.Close()

	for rows.Next() {
		todo := models.Todo{}

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Text, &todo.Time, &todo.Completed)

		if err != nil {
			log.Println("[ERROR] Not completed todo reading:", err)
			continue
		}

		todos = append(todos, todo)
	}

	return todos, nil
}
