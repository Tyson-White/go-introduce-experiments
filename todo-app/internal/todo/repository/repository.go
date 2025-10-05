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

	err := r.db.QueryRow(fmt.Sprintf("INSERT INTO todos (title, text) values ('%v', '%v')", data.Title, data.Text)).
		Scan(&createdTodo.Id, &createdTodo.Title, &createdTodo.Text, &createdTodo.Time, &createdTodo.Completed)

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
