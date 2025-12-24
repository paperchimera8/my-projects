package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type WorkTime struct {
	Store string
	Time  string
}

var stores = []WorkTime{
	{Store: "a", Time: "1"},
	{Store: "aa", Time: "11"},
	{Store: "a", Time: "1"},
	{Store: "aa", Time: "11"},
}

func (wt WorkTime) HandleTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	
	switch r.Method {
	case http.MethodGet:
		wt.handleGet(w, r)
	case http.MethodPost:
		wt.handlePost(w, r)
	case http.MethodPut:
		wt.handlePut(w, r)
	case http.MethodDelete:
		wt.handleDelete(w, r)
	}

}

func (wt WorkTime) handleGet(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(stores)
}

func (wt WorkTime) handlePost(w http.ResponseWriter, r *http.Request) {
	var workTime WorkTime
	err := json.NewDecoder(r.Body).Decode(&workTime)
	stores = append(stores, workTime)
	if err != nil {
		return 
	}
	json.NewEncoder(w).Encode(stores)
}

func (wt WorkTime) handlePut(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/stores/")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}
	userID--
	var newStore WorkTime
	err = json.NewDecoder(r.Body).Decode(&newStore)
	if err != nil {
		return
	}
	stores[userID] = newStore
	err = json.NewEncoder(w).Encode(stores)
}

func (wt WorkTime) handleDelete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/stores/")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}
	stores = append(stores[:userID-1], stores[userID:]...)
	json.NewEncoder(w).Encode(stores)
}

func main() {
	var wt WorkTime
	http.HandleFunc("/stores/", wt.HandleTasks)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}