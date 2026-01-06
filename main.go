package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"shop_list/internal/db"
	"shop_list/internal/handlers"
	"shop_list/internal/middleware"
	"shop_list/internal/repository"
	"shop_list/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.Connect()
	repo := repository.NewInMemoryShopRepository(db)
	service := services.NewShopServiceStruct(repo)
	handler := handlers.NewShopHandler(service)

	router := gin.Default()
	router.POST("/upload", UploadImage)
	router.Static("/images", "./upload")
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	router.Run(":8080")

	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)

	http.Handle("/stores/", middleware.AuthMiddleware(http.HandlerFunc(handler.HandleTasks)))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		return
	}
	os.MkdirAll("./upload", os.ModePerm)
	dst := filepath.Join("./upload", file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": "/images/" + file.Filename})
}
