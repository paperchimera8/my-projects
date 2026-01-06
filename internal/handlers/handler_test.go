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
)

func setupHandler() *ShopHandler {
	database := db.Connect()
	repo := repository.NewInMemoryShopRepository(database)
	service := services.NewShopServiceStruct(repo)
	return NewShopHandler(service)
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

	workTime := model.WorkTime{
		Name: "Тестовый магазин",
		Time: "10.00 - 22.00",
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

	workTime := model.WorkTime{
		ID:   1,
		Name: "Пятерочка Обновленная",
		Time: "8.00 - 23.00",
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

	createBody, _ := json.Marshal(model.WorkTime{
		Name: "Магазин для удаления",
		Time: "9.00 - 20.00",
	})
	createReq := httptest.NewRequest(http.MethodPost, "/stores/", bytes.NewReader(createBody))
	createRec := httptest.NewRecorder()
	handler.Create(createRec, createReq)

	workTime := model.WorkTime{
		ID: 3,
	}
	body, _ := json.Marshal(workTime)

	req := httptest.NewRequest(http.MethodDelete, "/stores/", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.Delete(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())
}
