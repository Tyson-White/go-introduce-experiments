package main

import (
	"arch-study/http"
	"fmt"
)

func main() {
	server := http.NewServer(http.AppHandler().Router)
	if err := server.Run(2000); err != nil {
		fmt.Println(err)
	}
}
