package handlers

import (
	"encoding/json"
	"net/http"
	"shop_list/internal/middleware"
	"shop_list/internal/model"
	"shop_list/internal/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type ShopHandler struct {
	service services.ShopService
}

func NewShopHandler(service services.ShopService) *ShopHandler {
	return &ShopHandler{
		service: service,
	}
}

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

type GoodHandler struct {
	service services.GoodService
}

func NewGoodHandler(service services.GoodService) *GoodHandler {
	return &GoodHandler{
		service: service,
	}
}
func (s ShopHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	stores, err := s.service.Get(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(stores)
}

func (s ShopHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var workTime model.Shop
	err := json.NewDecoder(r.Body).Decode(&workTime)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	layout := "15:04"

	openTime, err := time.Parse(layout, workTime.OpenTime)
	if err != nil {
		http.Error(w, "invalid open_time", http.StatusBadRequest)
		return
	}

	closeTime, err := time.Parse(layout, workTime.CloseTime)
	if err != nil {
		http.Error(w, "invalid close_time", http.StatusBadRequest)
		return
	}

	workTime.OpenTime = openTime.Format(layout)
	workTime.CloseTime = closeTime.Format(layout)

	err = s.service.Create(ctx, workTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s ShopHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var workTime model.Shop
	err := json.NewDecoder(r.Body).Decode(&workTime)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	layout := "15:04"

	openTime, err := time.Parse(layout, workTime.OpenTime)
	if err != nil {
		http.Error(w, "invalid open_time", http.StatusBadRequest)
		return
	}

	closeTime, err := time.Parse(layout, workTime.CloseTime)
	if err != nil {
		http.Error(w, "invalid close_time", http.StatusBadRequest)
		return
	}

	workTime.OpenTime = openTime.Format(layout)
	workTime.CloseTime = closeTime.Format(layout)

	err = s.service.Update(ctx, workTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s ShopHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var workTime model.Shop
	err := json.NewDecoder(r.Body).Decode(&workTime)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	idStr := workTime.ID.Hex()
	err = s.service.Delete(ctx, idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	user.Password = string(bytes)
	err = s.service.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user model.User
	var token string
	json.NewDecoder(r.Body).Decode(&user)
	findPassword, err := s.service.ComparePas(ctx, user)
	if err != nil {
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(findPassword), []byte(user.Password))
	if err != nil {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}

	userID, err := s.service.FindByID(ctx, user)
	if err != nil {
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}
	token, err = middleware.CreateToken(strconv.Itoa(userID))
	if err != nil {
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (s GoodHandler) UploadGoods(c *gin.Context) {
	ctx := c.Request.Context()
	var goods model.Goods
	err := c.ShouldBindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = s.service.CreateGoods(ctx, goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create goods"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Goods created successfully",
	})
}

func (s ShopHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		s.Get(w, r)
	case http.MethodPost:
		s.Create(w, r)
	case http.MethodPut:
		s.Update(w, r)
	case http.MethodDelete:
		s.Delete(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
