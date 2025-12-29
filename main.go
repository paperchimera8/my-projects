package main

import (
	"fmt"
	"shop_list/internal/handlers"
	"shop_list/internal/repository"
	"shop_list/internal/services"
	"net/http"
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
