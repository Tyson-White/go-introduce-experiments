package category

import (
	"db-study/pkg"
	"db-study/pkg/dto"
	"encoding/json"
	"errors"
	"net/http"
)

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	var body dto.DeleteCategoryBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, pkg.HttpError(err), http.StatusBadRequest)
		return
	}

	err := h.service.DeleteCategory(body.Category)

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
