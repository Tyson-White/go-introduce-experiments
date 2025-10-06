package httpserver

import (
	"fmt"
	"log"
	"net/http"
)

func StartHttpServer(handler http.Handler, port int) {

	log.Println("[INFO] Server listening on port:", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), handler); err != nil {
		log.Fatalln(err)
	}
}
