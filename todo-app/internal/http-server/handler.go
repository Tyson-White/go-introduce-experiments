package httpserver

import (
	"database/sql"

	todoHandler "db-study/internal/todo/handler"
	todoRepository "db-study/internal/todo/repository"
	todoService "db-study/internal/todo/service"

	"github.com/gorilla/mux"
)

type SubHandler interface {
	Handler(*mux.Router, any)
}

func RootHandler(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	todoR := todoRepository.New(db)
	todoS := todoService.New(todoR)

	todoHandler.Handler(r.PathPrefix("/todos").Subrouter(), todoS)

	return r
}
