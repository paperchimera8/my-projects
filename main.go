package main

import (
	"fmt"
	"net/http"
	"shop_list/internal/db"
	"shop_list/internal/handlers"
	"shop_list/internal/middleware"
	"shop_list/internal/repository"
	"shop_list/internal/services"
)

func main() {
	db := db.Connect()
	repo := repository.NewInMemoryShopRepository(db)
	service := services.NewShopServiceStruct(repo)
	handler := handlers.NewShopHandler(service)

	// Публичные endpoints (без аутентификации)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)

	http.Handle("/stores/", middleware.AuthMiddleware(http.HandlerFunc(handler.HandleTasks)))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
