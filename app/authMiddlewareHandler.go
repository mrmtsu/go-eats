package app

import (
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// _, err := r.Cookie("jwt")
		// if err != nil {
		// 	http.Error(w, "No admin cookie", http.StatusForbidden)
		// 	return
		// }
		// next.ServeHTTP(w, r)
		cookie, err := r.Cookie("jwt")
     if err != nil {
          log.Fatal("Cookie: ", err)
     }
		if _, err := ParseJwt(cookie.Value); err != nil {
			http.Error(w, "No admin cookie", http.StatusForbidden)
		}
		next.ServeHTTP(w, r)
	})
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}
