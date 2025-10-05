package todo

import (
	todo "db-study/internal/todo/service"
	"db-study/pkg"
	"db-study/pkg/dto"
	"encoding/json"
	"errors"
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

	return &h

}

func (h *TodoHandler) add(w http.ResponseWriter, r *http.Request) {

	var body dto.CreateTodo

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {

		http.Error(w, pkg.HttpError(err), http.StatusBadRequest)
		return
	}

	todo, err := h.service.AddTodo(body)

	if err != nil {
		if errors.Is(err, pkg.ErrValidation) {
			http.Error(w, pkg.HttpError(err), http.StatusBadRequest)
			return
		}

		http.Error(w, pkg.HttpError(err), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(todo)

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func (h *TodoHandler) allTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.Todos()

	if err != nil {

		http.Error(w, pkg.HttpError(err), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(todos)

	w.Write(data)

}
