package http

import (
	"fmt"

	"github.com/gorilla/mux"
)

type Handler interface {
	Router(r *mux.Router)
}

type RootHandler struct {
	Router *mux.Router
}

func NewRootHandler() *RootHandler {
	router := mux.NewRouter()

	return &RootHandler{
		Router: router,
	}

}

func (root *RootHandler) RegisterHandler(subHandler Handler, prefix string) *RootHandler {
	r := root.Router.PathPrefix(fmt.Sprintf("/%v", prefix)).Subrouter()
	subHandler.Router(r)

	return root
}
