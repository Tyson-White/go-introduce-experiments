package main

import (
	httpserver "db-study/internal/http-server"
	"db-study/internal/repository"
	"db-study/pkg"
)

func main() {
	repository := repository.ConnectDB(
		"postgres",
		"password",
		"postgres",
		5433,
	)

	handler := httpserver.RootHandler(repository.DB)

	go pkg.CleanFile("files/logs.txt")

	httpserver.StartHttpServer(handler, 8080)
}

/*

1. Сделать логгер запросов с выводом в .txt файл

*/
