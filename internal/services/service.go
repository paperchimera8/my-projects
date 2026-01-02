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
	Get() *sql.Rows
	Create(wt model.WorkTime) sql.Result
	Update(wt model.WorkTime) sql.Result
	Delete(id int) sql.Result
}

func (s ShopServiceStruct) Get() *sql.Rows {
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
