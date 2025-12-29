package repository

import (
	"shop_list/internal/model"
)

type ShopRepository interface {
	Get() []model.WorkTime
	Create(shop *model.WorkTime) *model.WorkTime
	Update(id int, shop *model.WorkTime)
	Delete(id int)
}

type InMemoryShopRepository struct {
	repo []model.WorkTime
}

func NewInMemoryShopRepository() *InMemoryShopRepository {
	return &InMemoryShopRepository{
		repo: []model.WorkTime{
			{Store: "a", Time: "1"},
			{Store: "aa", Time: "11"},
			{Store: "a", Time: "1"},
			{Store: "aa", Time: "11"},
		},
	}
}

func (r InMemoryShopRepository) Get() []model.WorkTime {
	return r.repo
}

func (r InMemoryShopRepository) Create(shop *model.WorkTime) *model.WorkTime {
	r.repo = append(r.repo, *shop)
	return shop
}

func (r *InMemoryShopRepository) Update(id int, shop *model.WorkTime) {
	r.repo[id] = *shop
}

func (r *InMemoryShopRepository) Delete(id int) {
	r.repo = append(r.repo[:id], r.repo[id+1:]...)
}
