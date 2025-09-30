package httpserver

import (
	"calc-app/calculator"

	"github.com/gorilla/mux"
)

func handler() *mux.Router {
	router := mux.NewRouter()

	calculator.CalculationsRouter(router)
	calculator.ResultsRouter(router)

	return router
}
