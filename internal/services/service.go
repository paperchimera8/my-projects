package services

import (
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
	Get() []model.WorkTime
	Create(shop *model.WorkTime) *model.WorkTime
	Update(id int, shop *model.WorkTime)
	Delete(id int)
}

func (s ShopServiceStruct) Get() []model.WorkTime {
	return s.repo.Get()
}

func (s *ShopServiceStruct) Create(shop *model.WorkTime) *model.WorkTime {
	// тут может быть бизнес-логика (валидация, проверки)
	return s.repo.Create(shop)
}

func (s *ShopServiceStruct) Update(id int, shop *model.WorkTime) {
	s.repo.Update(id, shop)
}

func (s *ShopServiceStruct) Delete(id int) {
	s.repo.Delete(id)
}
