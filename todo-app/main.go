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
		5432,
	)

	handler := httpserver.RootHandler(repository.DB)

	httpserver.StartHttpServer(handler, 8080)
}

/*

1. Привязка дел к категории по id

*/
