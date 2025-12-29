package internal

import (
	"internal/models"
)

type ShopServiceStruct struct {
	repo ShopRepository
}

func NewShopServiceStruct(repo ShopRepository) *ShopServiceStruct {
	return &ShopServiceStruct{repo: repo}
}

type ShopService interface {
	Get() []models.WorkTime
	Create(shop *models.WorkTime) *models.WorkTime
	Update(id int, shop *models.WorkTime)
	Delete(id int)
}

func (s ShopServiceStruct) Get() []models.WorkTime {
	return s.repo.Get()
}

func (s *ShopServiceStruct) Create(shop *models.WorkTime) *models.WorkTime {
	// тут может быть бизнес-логика (валидация, проверки)
	return s.repo.Create(shop)
}

func (s *ShopServiceStruct) Update(id int, shop *models.WorkTime) {
	s.repo.Update(id, shop)
}

func (s *ShopServiceStruct) Delete(id int) {
	s.repo.Delete(id)
}
