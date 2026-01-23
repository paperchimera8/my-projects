package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"shop_list/internal/db"
	"shop_list/internal/model"
	"shop_list/internal/repository"
	"shop_list/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setupHandler() *ShopHandler {
	database := db.Connect()
	repo := repository.NewShopRepositoryConstruct(database)
	service := services.NewShopServiceStruct(repo)
	return NewShopHandler(service)
}

func userHandler() *UserHandler {
	database := db.Connect()
	repo := repository.NewUserRepositoryConstruct(database)
	service := services.NewUserServiceStruct(repo)
	return NewUserHandler(service)
}
func TestGetHandler(t *testing.T) {
	handler := setupHandler()

	req := httptest.NewRequest(http.MethodGet, "/stores/", nil)
	rec := httptest.NewRecorder()

	handler.Get(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())
}

func TestCreateHandler(t *testing.T) {
	handler := setupHandler()

	workTime := model.Shop{
		Name:       "Тестовый магазин",
		OpenTime:   "09:00",
		CloseTime:  "18:00",
		DaysOfWeek: "пн-пт",
	}
	body, _ := json.Marshal(workTime)

	req := httptest.NewRequest(http.MethodPost, "/stores/", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())
}

func TestUpdateHandler(t *testing.T) {
	handler := setupHandler()

	workTime := model.Shop{
		Name:       "Тестовый магазин",
		OpenTime:   "09:00",
		CloseTime:  "18:00",
		DaysOfWeek: "пн-пт",
	}
	body, _ := json.Marshal(workTime)

	req := httptest.NewRequest(http.MethodPut, "/stores/", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.Update(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())
}

func TestDeleteHandler(t *testing.T) {
	handler := setupHandler()

	createBody, _ := json.Marshal(model.Shop{
		Name:       "Тестовый магазин",
		OpenTime:   "09:00",
		CloseTime:  "18:00",
		DaysOfWeek: "пн-пт",
	})
	createReq := httptest.NewRequest(http.MethodPost, "/stores/", bytes.NewReader(createBody))
	createRec := httptest.NewRecorder()
	handler.Create(createRec, createReq)

	objectID := primitive.NewObjectID()
	workTime := model.Shop{
		ID: objectID,
	}
	body, _ := json.Marshal(workTime)

	req := httptest.NewRequest(http.MethodDelete, "/stores/", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.Delete(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())
}
