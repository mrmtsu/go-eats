package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrmtsu/go-eats/domain"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := []domain.Article{}
	DB.Preload("User").Find(&articles)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	articleId := params["id"]

	article := domain.Article{}
	DB.Preload("User").Find(&article, articleId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func CreateAllArticles(w http.ResponseWriter, r *http.Request) {
	article := domain.Article{}
	json.NewDecoder(r.Body).Decode(&article)
	DB.Create(&article)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func UpdateArticles(w http.ResponseWriter, r *http.Request) {
	article := domain.Article{}
	json.NewDecoder(r.Body).Decode(&article)
	DB.Save(&article)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func DeleteArticles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	articleId := params["id"]
	DB.Delete(domain.Article{}, articleId)
	w.WriteHeader(http.StatusNoContent)
}
