package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mrmtsu/go-eats/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func sanityCheck() {
	godotenv.Load("envfiles/.env")

	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined ...")
	}
}

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	dbConnect()

	// User
	router.HandleFunc("/api/restaurant", GetAllRestaurant).Methods("GET")
	router.HandleFunc("/api/restaurant", CreaetRestaurant).Methods("POST")
	router.HandleFunc("/api/restaurant/{id}", GetRestaurant).Methods("GET")
	router.HandleFunc("/api/restaurant/{id}", UpdateRestaurant).Methods("PUT")
	router.HandleFunc("/api/restaurant/{id}", DeleteRestaurant).Methods("DELETE")

	// Article
	router.HandleFunc("/api/foods", GetAllFoods).Methods("GET")
	router.HandleFunc("/api/food", CreateAllFood).Methods("POST")
	router.HandleFunc("/api/food/{id}", GetFood).Methods("GET")
	router.HandleFunc("/api/food/{id}", UpdateFood).Methods("PUT")
	router.HandleFunc("/api/food/{id}", DeleteFood).Methods("DELETE")

	port := os.Getenv("SERVER_PORT")
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

var DB *gorm.DB

func dbConnect() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	dns := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8"
	database, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	DB = database

	database.AutoMigrate(&domain.User{}, &domain.Article{})
}
