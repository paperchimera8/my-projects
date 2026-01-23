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
	client := db.Connect()
	db.CreateIndexes(client) // Создаем индексы

	shopRepo := repository.NewShopRepositoryConstruct(client)
	shopService := services.NewShopServiceStruct(shopRepo)
	shopHandler := handlers.NewShopHandler(shopService)

	// users
	userRepo := repository.NewUserRepositoryConstruct(client)
	userService := services.NewUserServiceStruct(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// goods
	goodRepo := repository.NewGoodRepositoryConstruct(client)
	goodService := services.NewGoodServiceStruct(goodRepo)
	goodHandler := handlers.NewGoodHandler(goodService)

	router := gin.Default()

	// Загрузка файлов
	router.POST("/upload", UploadImage)
	router.Static("/images", "./upload")
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// API для магазинов (с JWT защитой)
	storesGroup := router.Group("/stores")
	storesGroup.Use(ginAuthMiddleware())
	{
		storesGroup.GET("/", ginHandler(shopHandler.Get))
		storesGroup.POST("/", ginHandler(shopHandler.Create))
		storesGroup.PUT("/", ginHandler(shopHandler.Update))
		storesGroup.DELETE("/", ginHandler(shopHandler.Delete))
	}

	// API для пользователей (публичные)
	router.POST("/register", ginHandler(userHandler.Register))
	router.POST("/login", ginHandler(userHandler.Login))

	// API для товаров
	router.POST("/goods", goodHandler.UploadGoods)

	fmt.Println("Сервер запущен на http://localhost:8080")
	router.Run(":8080")
}

// ginHandler конвертирует http.HandlerFunc в gin.HandlerFunc
func ginHandler(h func(http.ResponseWriter, *http.Request)) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c.Writer, c.Request)
	}
}

// ginAuthMiddleware конвертирует middleware в gin middleware
func ginAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authMiddleware := middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Next()
		}))
		authMiddleware.ServeHTTP(c.Writer, c.Request)
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
