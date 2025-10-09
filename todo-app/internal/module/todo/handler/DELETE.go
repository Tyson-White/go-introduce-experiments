package todo

import (
	"db-study/pkg"
	"db-study/pkg/dto"
	"encoding/json"
	"errors"
	"net/http"
)

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var body dto.DeleteTodoBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, pkg.HttpError(err), http.StatusBadRequest)
		return
	}

	err := h.service.DeleteTodo(body.Todo)

	if err != nil {

		if errors.Is(err, pkg.ErrNotFound) {
			http.Error(w, pkg.HttpError(err), http.StatusNotFound)
			return
		}

		http.Error(w, pkg.HttpError(err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
