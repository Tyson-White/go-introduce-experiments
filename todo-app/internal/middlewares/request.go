package middlewares

import (
	"db-study/pkg"
	"net/http"
	"sync"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
	Bytes  []byte
}

// Функция будет отрабатывать в хендлерах за место стандартных методов ResponseWriter
func (sr *StatusRecorder) WriteHeader(code int) {
	sr.Status = code
	sr.ResponseWriter.WriteHeader(code)
}

// Функция будет отрабатывать в хендлерах за место стандартных методов ResponseWriter
func (sr *StatusRecorder) Write(b []byte) (int, error) {
	if sr.Status == 0 {
		sr.Status = http.StatusOK // если WriteHeader не вызывали
	}
	n, err := sr.ResponseWriter.Write(b)
	sr.Bytes = b
	return n, err
}

var mtx sync.Mutex

// Обертка над ResponseWriter. Его методы перехватываются,
// чтобы производить манипуляции, после отработки запроса
func RequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &StatusRecorder{
			ResponseWriter: w,
		}

		// Надо узнать что делает ServeHTTP
		next.ServeHTTP(recorder, r)

		pkg.LogAndSave(recorder.Status, recorder.Bytes, r)

	})
}
