package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type userHandler struct {
}

var UserHandler = &userHandler{}

func (h *userHandler) Router(r *mux.Router) {
	r.Path("/greeting").Methods(http.MethodGet).HandlerFunc(h.printHelloHandler)

}

func (h *userHandler) printHelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!!!"))
}
