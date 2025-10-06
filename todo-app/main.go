package main

import (
	httpserver "db-study/internal/http-server"
	"db-study/internal/repository"
)

func main() {
	repository := repository.ConnectDB(
		"postgres",
		"password",
		"postgres",
		5433,
	)

	handler := httpserver.RootHandler(repository.DB)

	httpserver.StartHttpServer(handler, 8080)
}
