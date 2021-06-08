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
	s := router.PathPrefix("/api/auth").Subrouter().StrictSlash(false)

	dbConnect()

	// Auth
	router.HandleFunc("/api/register", Register).Methods("POST")
	router.HandleFunc("/api/login", Login).Methods("POST")

	// Middleware
	s.Handle("/user", IsAuthenticated(http.HandlerFunc(User))).Methods("GET")
	router.HandleFunc("/api/logout", Logout).Methods("GET")

	// User
	s.Handle("/users", IsAuthenticated(http.HandlerFunc(GetAllUsers))).Methods("GET")
	s.Handle("/users", IsAuthenticated(http.HandlerFunc(CreaetUsers))).Methods("POST")
	s.Handle("/users/{id}", IsAuthenticated(http.HandlerFunc(GetUsers))).Methods("GET")
	s.Handle("/users/{id}", IsAuthenticated(http.HandlerFunc(UpdateUsers))).Methods("PUT")
	s.Handle("/users/{id}", IsAuthenticated(http.HandlerFunc(DeleteUsers))).Methods("DELETE")

	// Article
	s.Handle("/articles", IsAuthenticated(http.HandlerFunc(GetAllArticles))).Methods("GET")
	s.Handle("/articles", IsAuthenticated(http.HandlerFunc(CreateAllArticles))).Methods("POST")
	s.Handle("/articles/{id}", IsAuthenticated(http.HandlerFunc(GetArticles))).Methods("GET")
	s.Handle("/articles/{id}", IsAuthenticated(http.HandlerFunc(UpdateArticles))).Methods("PUT")
	s.Handle("/articles/{id}", IsAuthenticated(http.HandlerFunc(DeleteArticles))).Methods("DELETE")

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

	database.AutoMigrate(&domain.User{}, &domain.Article{}, &domain.Comment{})
}
