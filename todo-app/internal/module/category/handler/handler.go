package category

import (
	category "db-study/internal/module/category/service"
	"net/http"

	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	service *category.CategoryService
}

func Handler(r *mux.Router, service *category.CategoryService) {

	h := CategoryHandler{
		service: service,
	}

	r.Path("/create").Methods(http.MethodPost).HandlerFunc(h.CreateCategory)
	r.Path("").Methods(http.MethodGet).HandlerFunc(h.CategoriesAll)
}
