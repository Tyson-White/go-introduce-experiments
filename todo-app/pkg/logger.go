package pkg

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var mtx sync.Mutex

func LogAndSave(status int, data []byte, r *http.Request) {

	var file *os.File

	mtx.Lock()
	defer func() {
		file.Close()

		mtx.Unlock()
	}()

	lType := "[REQUEST]"
	l := fmt.Sprintf("%v %d %s %v", lType, status, r.Method, r.URL)

	if status >= 400 {
		lType = "[REQUEST ERROR]"
	}

	log.Println(l)
	log.Printf("[RESPONSE BODY] %s", string(data))

	file, err := os.OpenFile("./files/logs.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Println("[ERROR] Log:", err)
		return
	}

	_, err = file.WriteString(fmt.Sprintln(time.Now(), l, "\n", string(data), "\n"))

	if err != nil {
		log.Println("[ERROR] Log:", err)
		return
	}

}
