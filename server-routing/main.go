package main

import (
	"arch-study/http"
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		server := http.NewServer(http.AppHandler().Router)
		if err := server.Run(2000); err != nil {
			fmt.Println(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		testServer := http.NewServer(http.TestHandler().Router)
		if err := testServer.Run(1000); err != nil {
			fmt.Println(err)
		}
	}()

	wg.Wait()
}
