package main

import (
	"fmt"
	"net/http"
	"internal/handlers"
	"internal/repository"
	"internal/services"
)

func main() {
	repo := repository.NewInMemoryShopRepository()
	service := services.NewShopServiceStruct(repo)
	handler := handlers.NewShopHandler(service)
	
	http.HandleFunc("/stores/", handler.Get)
	http.HandleFunc("/stores", handler.Get)
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
