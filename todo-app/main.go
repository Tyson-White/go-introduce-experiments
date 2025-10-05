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
	)

	handler := httpserver.RootHandler(repository.DB)

	httpserver.StartHttpServer(handler)
}
