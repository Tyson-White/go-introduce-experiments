package category

import (
	"db-study/pkg"
	"encoding/json"
	"net/http"
)

func (h *CategoryHandler) CategoriesAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.CategoriesAll()

	if err != nil {
		http.Error(w, pkg.HttpError(err), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(categories)

	if err != nil {
		http.Error(w, pkg.HttpError(err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(data)
}
