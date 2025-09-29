package http

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var mtx = sync.Mutex{}

type Server struct {
	router *mux.Router
}

var usedPorts = map[int]bool{}

func NewServer(router *mux.Router) *Server {

	return &Server{
		router: router,
	}
}

func (s *Server) Run(port int) error {
	mtx.Lock()

	if _, ok := usedPorts[port]; ok {
		mtx.Unlock()
		return errors.New(fmt.Sprintf("port %d is already used", port))
	}

	usedPorts[port] = true
	mtx.Unlock()

	fmt.Println("Server listening in port:", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s.router)
}
