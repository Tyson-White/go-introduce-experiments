package httpserver

import (
	"db-study/internal/middlewares"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func StartHttpServer(handler http.Handler, port int) {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	log.Println("[INFO] Server listening on port:", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), c.Handler(middlewares.RequestLog(handler))); err != nil {
		log.Fatalln(err)
	}
}
