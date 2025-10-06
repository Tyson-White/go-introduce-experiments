package category

import (
	"db-study/pkg"
	"db-study/pkg/dto"
	"encoding/json"
	"errors"
	"net/http"
)

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var body dto.CreateCategory

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, pkg.HttpError(err), http.StatusBadRequest)
		return
	}

	cat, err := h.service.AddCategory(body.Name)
	if err != nil {
		if errors.Is(err, pkg.ErrValidation) {
			http.Error(w, pkg.HttpError(err), http.StatusBadRequest)

			return
		}

		http.Error(w, pkg.HttpError(err), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(cat)

	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}
