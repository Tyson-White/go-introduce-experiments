package calculator

import (
	"calc-app/calculator/operations"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func ResultsRouter(r *mux.Router) {
	subr := r.PathPrefix("/results").Subrouter()

	subr.Path("").Methods(http.MethodGet).HandlerFunc(historyAllHandler)
	subr.Path("/{type}").Methods(http.MethodGet).HandlerFunc(historyByTypeHandler)

}

func historyAllHandler(w http.ResponseWriter, r *http.Request) {
	history := operations.Operations()

	data, err := json.Marshal(history)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func historyByTypeHandler(w http.ResponseWriter, r *http.Request) {
	t := mux.Vars(r)["type"]

	operations := operations.OperationsByType(t)

	data, err := json.Marshal(operations)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)

}
