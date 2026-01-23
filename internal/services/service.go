package services

import (
	"context"
	"fmt"
	"shop_list/internal/model"
	"shop_list/internal/repository"
)

type ShopServiceStruct struct {
	repo repository.ShopRepository
}

func NewShopServiceStruct(repo repository.ShopRepository) *ShopServiceStruct {
	return &ShopServiceStruct{repo: repo}
}

type UserServiceStruct struct {
	repo repository.UserRepository
}

func NewUserServiceStruct(repo repository.UserRepository) *UserServiceStruct {
	return &UserServiceStruct{repo: repo}
}

type GoodServiceStruct struct {
	repo repository.GoodRepository
}

func NewGoodServiceStruct(repo repository.GoodRepository) *GoodServiceStruct {
	return &GoodServiceStruct{repo: repo}
}

type ShopService interface {
	Get(ctx context.Context) ([]model.Shop, error)
	Create(ctx context.Context, wt model.Shop) error
	Update(ctx context.Context, wt model.Shop) error
	Delete(ctx context.Context, id string) error
}

type UserService interface {
	CreateUser(ctx context.Context, u model.User) error
	ComparePas(ctx context.Context, u model.User) (string, error)
	FindByID(ctx context.Context, u model.User) (int, error)
}

type GoodService interface {
	CreateGoods(ctx context.Context, g model.Goods) error
}

func (s *ShopServiceStruct) Get(ctx context.Context) ([]model.Shop, error) {
	res, err := s.repo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("get shops: %w", err)
	}
	return res, nil
}

func (s *ShopServiceStruct) Create(ctx context.Context, wt model.Shop) error {
	// тут может быть бизнес-логика (валидация, проверки)
	err := s.repo.Create(ctx, wt)
	if err != nil {
		return fmt.Errorf("create shops: %w", err)
	}
	return nil
}

func (s *ShopServiceStruct) Update(ctx context.Context, wt model.Shop) error {
	err := s.repo.Update(ctx, wt)
	if err != nil {
		return fmt.Errorf("update shops: %w", err)
	}
	return nil
}

func (s *ShopServiceStruct) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("delete shops: %w", err)
	}
	return nil
}

func (s *UserServiceStruct) CreateUser(ctx context.Context, u model.User) error {
	err := s.repo.CreateUser(ctx, u)
	if err != nil {
		return fmt.Errorf("create user error: %w", err)
	}
	return nil
}

func (s *UserServiceStruct) ComparePas(ctx context.Context, u model.User) (string, error) {
	res, err := s.repo.ComparePas(ctx, u)
	if err != nil {
		return "", fmt.Errorf("compare password error: %w", err)
	}
	return res, nil
}

func (s *UserServiceStruct) FindByID(ctx context.Context, u model.User) (int, error) {
	res, err := s.repo.FindByID(ctx, u)
	if err != nil {
		return 0, fmt.Errorf("findbyid error: %w", err)
	}
	return res, nil
}

func (s *GoodServiceStruct) CreateGoods(ctx context.Context, g model.Goods) error {
	err := s.repo.CreateGoods(ctx, g)
	if err != nil {
		return fmt.Errorf("create goods error: %w", err)
	}
	return nil
}
