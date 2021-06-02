package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrmtsu/go-eats/domain"
)

func GetAllRestaurant(w http.ResponseWriter, r *http.Request) {
	data := DB.Preload("Article.User").Begin()
	users := []domain.User{}
	data.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetRestaurant(w http.ResponseWriter, r *http.Request) {
	data := DB.Preload("Article.User").Begin()
	params := mux.Vars(r)
	userId := params["id"]

	user := domain.User{}
	data.Find(&user, userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreaetRestaurant(w http.ResponseWriter, r *http.Request) {
	user := domain.User{}
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
	user := domain.User{}
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	DB.Delete(domain.User{}, userId)
	w.WriteHeader(http.StatusNoContent)
}
