package main

import (
	"fmt"
	"net/http"
	"internal/handler"
)

func main() {
	http.HandleFunc("/stores/", handler.HandleTasks)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
