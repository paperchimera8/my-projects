package repository

import (
	"database/sql"
	"shop_list/internal/model"
)

type ShopRepository interface {
	Get() ([]model.WorkTime, error)
	Create(wt model.WorkTime) sql.Result
	Update(wt model.WorkTime) sql.Result
	Delete(id int) sql.Result
	CreateUser(u model.User)
	ComparePas(query string, u model.User) string
	FindID(query string, u model.User) uint
}

type InMemoryShopRepository struct {
	repo *sql.DB
}

func NewInMemoryShopRepository(repo *sql.DB) *InMemoryShopRepository {
	return &InMemoryShopRepository{
		repo: repo,
	}
}

func (r InMemoryShopRepository) Get() ([]model.WorkTime, error) {
	rows, err := r.repo.Query("SELECT * FROM Shops")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shops []model.WorkTime
	for rows.Next() {
		var shop model.WorkTime
		err := rows.Scan(&shop.ID, &shop.Name, &shop.Time)
		if err != nil {
			return nil, err
		}
		shops = append(shops, shop)
	}
	return shops, nil
}

func (r InMemoryShopRepository) Create(wt model.WorkTime) sql.Result {
	result, err := r.repo.Exec("INSERT INTO Shops (name, time) VALUES ($1, $2)", wt.Name, wt.Time)
	if err != nil {
		panic(err)
	}
	return result
}

func (r *InMemoryShopRepository) Update(wt model.WorkTime) sql.Result {
	query := "UPDATE Shops SET name = $1, time = $2 WHERE id = $3"
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

func (r *InMemoryShopRepository) CreateUser(u model.User) {
	_, err := r.repo.Exec("INSERT INTO Users (username, password) VALUES ($1, $2) RETURNING id", u.Username, u.Password)
	if err != nil {
		panic(err)
	}
}

func (r *InMemoryShopRepository) ComparePas(query string, u model.User) string {
	var pas string
	pass := r.repo.QueryRow(query, u.Username)
	pass.Scan(&pas)
	return pas
}

func (r *InMemoryShopRepository) FindID(query string, u model.User) uint {
	var pas uint
	pass := r.repo.QueryRow(query, u.Username)
	pass.Scan(&pas)
	return pas
}
