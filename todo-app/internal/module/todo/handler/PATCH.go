package todo

import (
	"db-study/pkg"
	"db-study/pkg/dto"
	"encoding/json"
	"net/http"
)

func (h *TodoHandler) markAsCompleted(w http.ResponseWriter, r *http.Request) {
	var body dto.MarkTodo

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, pkg.HttpError(err), http.StatusBadRequest)
		return
	}

	todo, err := h.service.MarkAsCompleted(body.Id)

	if err != nil {
		http.Error(w, pkg.HttpError(err), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(todo)

	if err != nil {
		http.Error(w, pkg.HttpError(err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(data)
}
