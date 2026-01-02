package repository

import (
	"database/sql"
	"shop_list/internal/model"
)

type ShopRepository interface {
	Get() *sql.Rows
	Create(wt model.WorkTime) sql.Result
	Update(wt model.WorkTime) sql.Result
	Delete(id int) sql.Result
}

type InMemoryShopRepository struct {
	repo *sql.DB
}

func NewInMemoryShopRepository(repo *sql.DB) *InMemoryShopRepository {
	return &InMemoryShopRepository{
		repo: repo,
	}
}

func (r InMemoryShopRepository) Get() *sql.Rows {
	rows, err := r.repo.Query("SELECT * FROM Shops")
	if err != nil {
		panic(err)
	}
	return rows
}

func (r InMemoryShopRepository) Create(wt model.WorkTime) sql.Result {
	result, err := r.repo.Exec("INSERT INTO Shops (id, name, time) VALUES ($1, $2, $3)", wt.ID, wt.Name, wt.Time)
	if err != nil {
		panic(err)
	}
	return result
}

func (r *InMemoryShopRepository) Update(wt model.WorkTime) sql.Result {
	query := "UPDATE Shops SET (shopName, shopTime) VALUES ($1, $2) WHERE id = $3"
	result, err := r.repo.Exec(query, wt.Name, wt.Time, wt.ID)
	if err != nil {
		panic(err)
	}
	return result
}

func (r *InMemoryShopRepository) Delete(id int) sql.Result {
	query := "DELETE FROM Shops WHERE ID = $1"
	result, err := r.repo.Exec(query, id)
	if err != nil {
		panic(err)
	}
	return result
}
