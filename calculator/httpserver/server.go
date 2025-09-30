package httpserver

import (
	"fmt"
	"net/http"
)

func RunServer(port int) {
	handler := handler()

	fmt.Printf("Server listening on port: %v \n", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), handler); err != nil {
		fmt.Println("Server shutdown:", err)
	}
}
