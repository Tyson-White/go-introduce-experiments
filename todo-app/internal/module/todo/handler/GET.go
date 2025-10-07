package todo

import (
	"db-study/pkg"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *TodoHandler) allTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.Todos()

	if err != nil {

		http.Error(w, pkg.HttpError(err), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(todos)

	w.WriteHeader(http.StatusOK)
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

	w.WriteHeader(http.StatusOK)
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

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *TodoHandler) TodosByCategory(w http.ResponseWriter, r *http.Request) {
	catId := r.URL.Query().Get("category")
	id, err := strconv.Atoi(catId)

	if err != nil {
		http.Error(w, pkg.HttpError(err), http.StatusBadRequest)
		return
	}

	todos, err := h.service.TodoByCategory(id)

	if err != nil {
		http.Error(w, pkg.HttpError(err), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(todos)

	w.WriteHeader(http.StatusOK)
	w.Write(data)

}
