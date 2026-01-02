package handlers

import (
	"encoding/json"
	"net/http"
	"shop_list/internal/model"
	"shop_list/internal/services"
)

type ShopHandler struct {
	Service services.ShopService
}

func NewShopHandler(service services.ShopService) *ShopHandler {
	return &ShopHandler{
		Service: service,
	}
}

func (s ShopHandler) Get(w http.ResponseWriter, r *http.Request) {
	stores := s.Service.Get()
	json.NewEncoder(w).Encode(stores)
}

func (s ShopHandler) Create(w http.ResponseWriter, r *http.Request) {
	var workTime model.WorkTime
	err := json.NewDecoder(r.Body).Decode(&workTime)
	if err != nil {
		return
	}
	stores := s.Service.Create(workTime)
	json.NewEncoder(w).Encode(stores)
}

func (s ShopHandler) Update(w http.ResponseWriter, r *http.Request) {
	var workTime model.WorkTime
	err := json.NewDecoder(r.Body).Decode(&workTime)
	if err != nil {
		return
	}
	store := s.Service.Update(workTime)
	err = json.NewEncoder(w).Encode(store)
}

func (s ShopHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var workTime model.WorkTime
	err := json.NewDecoder(r.Body).Decode(&workTime)
	if err != nil {
		return
	}
	var id int = workTime.ID
	store := s.Service.Delete(id)
	err = json.NewEncoder(w).Encode(store)
	if err != nil {
		return
	}
}

func (s *ShopHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
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
