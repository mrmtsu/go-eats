package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mrmtsu/go-eats/domain"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	if data["password"] != data["password_confirm"] {
		http.NotFound(w, r)
		return
	}

	user := domain.User{
		Name:  data["name"],
		Email: data["email"],
	}

	user.SetPassword(data["password"])

	DB.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	var user domain.User

	DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		http.NotFound(w, r)
		return
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	if err != nil {
		w.WriteHeader(400)
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	w.Header().Set("Content-Type", "apllication/json")
	json.NewEncoder(w).Encode(user)
}

type Claims struct {
	jwt.StandardClaims
}

func User(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("jwt")
	id, _ := ParseJwt(cookie.Value)

	var user domain.User
	DB.Where("id = ?", id).First(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
}
