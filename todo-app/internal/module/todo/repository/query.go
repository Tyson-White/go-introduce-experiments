package repository

import (
	"db-study/pkg/models"
	"log"
)

func (r *TodoRepository) Todos() ([]models.Todo, error) {

	todos := []models.Todo{}

	rows, err := r.db.Queryx("SELECT * FROM todos")

	if err != nil {
		return todos, err
	}

	defer rows.Close()

	for rows.Next() {

		todo := models.Todo{}

		err := rows.StructScan(&todo)

		if err != nil {
			log.Println(err)
			continue
		}

		todos = append(todos, todo)
	}

	return todos, nil

}

func (r *TodoRepository) TodosNotCompleted() ([]models.Todo, error) {
	var todos = []models.Todo{}

	rows, err := r.db.Queryx("SELECT * FROM todos WHERE completed=false")

	if err != nil {
		return todos, err
	}

	defer rows.Close()

	for rows.Next() {
		todo := models.Todo{}

		err := rows.StructScan(&todo)

		if err != nil {
			log.Println("[ERROR] Not completed todo reading:", err)
			continue
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *TodoRepository) TodosCompleted() ([]models.Todo, error) {
	var todos = []models.Todo{}

	rows, err := r.db.Queryx("SELECT * FROM todos WHERE completed=true")

	if err != nil {
		return todos, err
	}

	defer rows.Close()

	for rows.Next() {
		todo := models.Todo{}

		err := rows.StructScan(&todo)

		if err != nil {
			log.Println("[ERROR] Not completed todo reading:", err)
			continue
		}

		todos = append(todos, todo)
	}

	return todos, nil
}
