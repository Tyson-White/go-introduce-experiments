package httpserver

import (
	todoHandler "db-study/internal/module/todo/handler"
	todoRepository "db-study/internal/module/todo/repository"
	todoService "db-study/internal/module/todo/service"

	categoryHandler "db-study/internal/module/category/handler"
	categoryRepository "db-study/internal/module/category/repository"
	categoryService "db-study/internal/module/category/service"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type SubHandler interface {
	Handler(*mux.Router, any)
}

func RootHandler(db *sqlx.DB) *mux.Router {
	r := mux.NewRouter()

	todoRepo := todoRepository.New(db)
	todoS := todoService.New(todoRepo)

	catRepo := categoryRepository.New(db)
	catS := categoryService.New(catRepo)

	todoHandler.Handler(r.PathPrefix("/todos").Subrouter(), todoS)
	categoryHandler.Handler(r.PathPrefix("/categories").Subrouter(), catS)

	return r
}
