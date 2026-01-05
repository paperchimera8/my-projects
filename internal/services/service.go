package services

import (
	"database/sql"
	"shop_list/internal/model"
	"shop_list/internal/repository"
)

type ShopServiceStruct struct {
	repo repository.ShopRepository
}

func NewShopServiceStruct(repo repository.ShopRepository) *ShopServiceStruct {
	return &ShopServiceStruct{repo: repo}
}

type ShopService interface {
	Get() ([]model.WorkTime, error)
	Create(wt model.WorkTime) sql.Result
	Update(wt model.WorkTime) sql.Result
	Delete(id int) sql.Result
	CreateUser(u model.User)
	ComparePas(query string, u model.User) string
	FindID(query string, u model.User) uint
}

func (s ShopServiceStruct) Get() ([]model.WorkTime, error) {
	return s.repo.Get()
}

func (s *ShopServiceStruct) Create(wt model.WorkTime) sql.Result {
	// тут может быть бизнес-логика (валидация, проверки)
	return s.repo.Create(wt)
}

func (s *ShopServiceStruct) Update(wt model.WorkTime) sql.Result {
	return s.repo.Update(wt)
}

func (s *ShopServiceStruct) Delete(id int) sql.Result {
	return s.repo.Delete(id)
}

func (s *ShopServiceStruct) CreateUser(u model.User) {
	s.repo.CreateUser(u)
}

func (s *ShopServiceStruct) ComparePas(query string, u model.User) string {
	return s.repo.ComparePas(query, u)
}

func (s *ShopServiceStruct) FindID(query string, u model.User) uint {
	return s.repo.FindID(query, u)
}
