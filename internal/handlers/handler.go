package handlers

import (
	"encoding/json"
	"net/http"
	"shop_list/internal/middleware"
	"shop_list/internal/model"
	"shop_list/internal/services"

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

func (s ShopHandler) Get(w http.ResponseWriter, r *http.Request) {
	stores, err := s.service.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(stores)
}

func (s ShopHandler) Create(w http.ResponseWriter, r *http.Request) {
	var workTime model.WorkTime
	err := json.NewDecoder(r.Body).Decode(&workTime)
	if err != nil {
		return
	}
	stores := s.service.Create(workTime)
	json.NewEncoder(w).Encode(stores)
}

func (s ShopHandler) Update(w http.ResponseWriter, r *http.Request) {
	var workTime model.WorkTime
	err := json.NewDecoder(r.Body).Decode(&workTime)
	if err != nil {
		return
	}
	store := s.service.Update(workTime)
	err = json.NewEncoder(w).Encode(store)
}

func (s ShopHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var workTime model.WorkTime
	err := json.NewDecoder(r.Body).Decode(&workTime)
	if err != nil {
		return
	}
	var id int = workTime.ID
	store := s.service.Delete(id)
	err = json.NewEncoder(w).Encode(store)
	if err != nil {
		return
	}
}

func (s ShopHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user.Password = string(bytes)
	s.service.CreateUser(user)
}

func (s *ShopHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var token string
	json.NewDecoder(r.Body).Decode(&user)
	query := "SELECT Password FROM Users WHERE username = $1"
	queryID := "SELECT ID FROM Users WHERE username = $1"
	findPassword := s.service.ComparePas(query, user)
	err := bcrypt.CompareHashAndPassword([]byte(findPassword), []byte(user.Password))
	if err != nil {
		return
	}
	userID := s.service.FindID(queryID, user)
	token, err = middleware.CreateToken(userID)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})

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
