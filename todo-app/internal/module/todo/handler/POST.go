package todo

import (
	"db-study/pkg"
	"db-study/pkg/dto"
	"encoding/json"
	"errors"
	"net/http"
)

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
