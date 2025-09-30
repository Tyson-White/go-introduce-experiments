package calculator

import (
	"calc-app/calculator/operations"
	"calc-app/dto"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CalculationsRouter(r *mux.Router) {

	// calc router
	subr := r.PathPrefix("/calc").Subrouter()

	subr.Path("/summary").Methods(http.MethodPost).HandlerFunc(calculateSumHandler)
	subr.Path("/difference").Methods(http.MethodPost).HandlerFunc(calculateDiffHandler)
	subr.Path("/multiple").Methods(http.MethodPost).HandlerFunc(calculateMultipleHandler)
	subr.Path("/divide").Methods(http.MethodPost).HandlerFunc(calculateDivideHandler)

}

// Возвращает результат сложения или ошибку.
func calculateSumHandler(w http.ResponseWriter, r *http.Request) {

	var body dto.OperationBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {

		http.Error(w, dto.NewErrorObj(err), http.StatusBadRequest)
		return
	}

	result := operations.SumOfAB(body.A, body.B)

	data, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)

}

// Возвращает результат вычетания или ошибку
func calculateDiffHandler(w http.ResponseWriter, r *http.Request) {
	var body dto.OperationBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {

		http.Error(w, dto.NewErrorObj(err), http.StatusBadRequest)
		return
	}

	result := operations.DifferenceOfAB(body.A, body.B)

	data, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)

}

func calculateMultipleHandler(w http.ResponseWriter, r *http.Request) {
	var body dto.OperationBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, dto.NewErrorObj(err), http.StatusBadRequest)
		return
	}

	result := operations.MultipleAB(body.A, body.B)

	data, err := json.Marshal(result)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)

}

func calculateDivideHandler(w http.ResponseWriter, r *http.Request) {

	var body dto.OperationBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, dto.NewErrorObj(err), http.StatusBadRequest)
		return
	}

	result := operations.DivideAB(body.A, body.B)

	data, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
