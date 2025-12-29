package internal

import (
	"internal/models"
)

type ShopRepository interface {
	Get() []models.WorkTime
	Create(shop *models.WorkTime) *models.WorkTime
	Update(id int, shop *models.WorkTime)
	Delete(id int)
}

type InMemoryShopRepository struct {
	repo []models.WorkTime
}

func NewInMemoryShopRepository() *InMemoryShopRepository {
	return &InMemoryShopRepository{
		repo: []models.WorkTime{
			{Store: "a", Time: "1"},
			{Store: "aa", Time: "11"},
			{Store: "a", Time: "1"},
			{Store: "aa", Time: "11"},
		},
	}
}

func (r InMemoryShopRepository) Get() []models.WorkTime {
	return r.repo
}

func (r InMemoryShopRepository) Create(shop *models.WorkTime) *models.WorkTime {
	r.repo = append(r.repo, *shop)
	return r.repo
}

func (r InMemoryShopRepository) Update(id int, shop *models.WorkTime) {
	r.repo[id] = shop
}

func (r InMemoryShopRepository) Delete(id int) {
	r.repo = append(r.repo[:id], r.repo[id+1:]...)
}
