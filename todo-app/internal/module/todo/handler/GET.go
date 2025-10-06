package todo

import (
	"db-study/pkg"
	"encoding/json"
	"net/http"
)

func (h *TodoHandler) allTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.Todos()

	if err != nil {

		http.Error(w, pkg.HttpError(err), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(todos)

	w.Write(data)

}

func (h *TodoHandler) TodosNotCompleted(w http.ResponseWriter, r *http.Request) {

	todos, err := h.service.TodosNotCompleted()

	if err != nil {
		http.Error(w, pkg.HttpError(err), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(todos)

	if err != nil {
		http.Error(w, pkg.HttpError(err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(data)
}

func (h *TodoHandler) TodosCompleted(w http.ResponseWriter, r *http.Request) {

	todos, err := h.service.TodosCompleted()

	if err != nil {
		http.Error(w, pkg.HttpError(err), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(todos)

	if err != nil {
		http.Error(w, pkg.HttpError(err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(data)
}
