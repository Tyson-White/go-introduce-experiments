package todo

import (
	todo "db-study/internal/module/todo/service"
	"net/http"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	service *todo.TodoService
}

func Handler(r *mux.Router, service *todo.TodoService) *TodoHandler {
	h := TodoHandler{
		service: service,
	}

	r.Path("/create").Methods(http.MethodPost).HandlerFunc(h.add)
	r.Path("").Methods(http.MethodGet).HandlerFunc(h.allTodos)
	r.Path("/complete").Methods(http.MethodPatch).HandlerFunc(h.markAsCompleted)
	r.Path("/not-completed").Methods(http.MethodGet).HandlerFunc(h.TodosNotCompleted)
	r.Path("/completed").Methods(http.MethodGet).HandlerFunc(h.TodosCompleted)

	return &h

}
