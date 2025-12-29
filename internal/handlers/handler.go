package handlers

import (
	"encoding/json"
	"internal/model"
	"internal/services"
	"net/http"
	"strconv"
	"strings"
)

type ShopHandlerInterface interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	s := ShopHandler{}
	switch r.Method {
	case http.MethodGet:
		s.Get(w, r)
	case http.MethodPost:
		s.Create(w, r)
	case http.MethodPut:
		s.Update(w, r)
	case http.MethodDelete:
		s.Delete(w, r)
	}

}

type ShopHandler struct {
	service services.ShopService
}

func NewShopHandler(service services.ShopService) *ShopHandler {
	return &ShopHandler{
		service: service,
	}
}

func (s ShopHandler) Get(w http.ResponseWriter, r *http.Request) {
	stores := s.service.Get()
	json.NewEncoder(w).Encode(stores)
}

func (s ShopHandler) Create(w http.ResponseWriter, r *http.Request) {
	var workTime model.WorkTime
	err := json.NewDecoder(r.Body).Decode(&workTime)
	if err != nil {
		return
	}
	stores := s.service.Create(&workTime)
	json.NewEncoder(w).Encode(stores)
}

func (s ShopHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/stores/")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}
	userID--
	var newStore model.WorkTime
	err = json.NewDecoder(r.Body).Decode(&newStore)
	if err != nil {
		return
	}
	s.service.Update(userID, &newStore)
	err = json.NewEncoder(w).Encode(s.service.Get())
}

func (s ShopHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/stores/")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}
	s.service.Delete(userID)
	err = json.NewEncoder(w).Encode(s.service.Get())
}
