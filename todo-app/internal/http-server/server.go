package httpserver

import (
	"log"
	"net/http"
)

func StartHttpServer(handler http.Handler) {

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalln(err)
	}
}
