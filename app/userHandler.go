package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrmtsu/go-eats/domain"
)

func GetAllRestaurant (w http.ResponseWriter, r *http.Request) {
	users := domain.User{}
	DB.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&users)
}

func GetRestaurant (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)
	userId := params["id"]

	user := domain.User{}
	DB.Find(&user, userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

func CreaetRestaurant (w http.ResponseWriter, r *http.Request) {
	user := domain.User{}
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
}
