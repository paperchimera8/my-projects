package repository

import (
	"context"
	"fmt"
	"shop_list/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShopRepository interface {
	Get(ctx context.Context) ([]model.Shop, error)
	Create(ctx context.Context, wt model.Shop) error
	Update(ctx context.Context, wt model.Shop) error
	Delete(ctx context.Context, id string) error
}

type UserRepository interface {
	CreateUser(ctx context.Context, u model.User) error
	ComparePas(ctx context.Context, u model.User) (string, error)
	FindByID(ctx context.Context, u model.User) (int, error)
}

type GoodRepository interface {
	CreateGoods(ctx context.Context, g model.Goods) error
}

type ShopRepositoryConstruct struct {
	repo *mongo.Collection
}

func NewShopRepositoryConstruct(client *mongo.Client) *ShopRepositoryConstruct {
	return &ShopRepositoryConstruct{
		repo: client.Database("mydb").Collection("shops"),
	}
}

type UserRepositoryConstruct struct {
	repo *mongo.Collection
}

func NewUserRepositoryConstruct(client *mongo.Client) *UserRepositoryConstruct {
	return &UserRepositoryConstruct{
		repo: client.
			Database("mydb").
			Collection("users"),
	}
}

type GoodRepositoryConstruct struct {
	repo *mongo.Collection
}

func NewGoodRepositoryConstruct(client *mongo.Client) *GoodRepositoryConstruct {
	return &GoodRepositoryConstruct{
		repo: client.
			Database("mydb").
			Collection("goods"),
	}
}

func (r *ShopRepositoryConstruct) Get(ctx context.Context) ([]model.Shop, error) {
	cursor, err := r.repo.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("get shops: %w", err)
	}
	defer cursor.Close(ctx)

	var shops []model.Shop
	err = cursor.All(ctx, &shops)
	if err != nil {
		return nil, fmt.Errorf("get shops parsing: %w", err)
	}
	return shops, nil
}

func (r *ShopRepositoryConstruct) Create(ctx context.Context, wt model.Shop) error {
	_, err := r.repo.InsertOne(ctx, wt)
	if err != nil {
		return fmt.Errorf("create shops: %w", err)
	}
	return nil
}

func (r *ShopRepositoryConstruct) Update(ctx context.Context, wt model.Shop) error {
	filter := bson.D{{Key: "_id", Value: wt.ID}}
	update := bson.M{
		"$set": bson.M{
			"name":       wt.Name,
			"opentime":   wt.OpenTime,
			"closetime":  wt.CloseTime,
			"daysofweek": wt.DaysOfWeek,
		},
	}
	_, err := r.repo.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update shops: %w", err)
	}
	return nil
}

func (r *ShopRepositoryConstruct) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": objectID}
	_, err = r.repo.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete shops: %w", err)
	}
	return nil
}

func (r *UserRepositoryConstruct) CreateUser(ctx context.Context, u model.User) error {
	_, err := r.repo.InsertOne(ctx, bson.M{
		"username": u.Username,
		"password": u.Password,
	})
	if err != nil {
		return fmt.Errorf("create user error: %w", err)
	}
	return nil
}

func (r *UserRepositoryConstruct) ComparePas(ctx context.Context, u model.User) (string, error) {
	var pass string
	result := r.repo.FindOne(ctx,
		bson.M{
			"username": u.Username,
		},
		options.FindOne().SetProjection(bson.M{
			"password": 1,
			"ID":       0, // обязательно, иначе вернётся _id
		}),
	)
	err := result.Decode(&pass)
	if err != nil {
		return "", fmt.Errorf("compare password error: %w", err)
	}
	return pass, nil
}

func (r *UserRepositoryConstruct) FindByID(ctx context.Context, u model.User) (int, error) {
	var result int
	err := r.repo.FindOne(ctx,
		bson.M{"username": u.Username},
		options.FindOne().SetProjection(bson.M{
			"_id":      1,
			"password": 0,
		}),
	).Decode(&result)
	if err != nil {
		return 0, fmt.Errorf("findbyid error: %w", err)
	}
	return result, nil
}

func (r *GoodRepositoryConstruct) CreateGoods(ctx context.Context, g model.Goods) error {
	_, err := r.repo.InsertOne(ctx, g)
	if err != nil {
		return fmt.Errorf("create goods eror: %w", err)
	}
	return nil
}
