package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TokenHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if username != "myusername" || password != "mypassword" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"exp":  time.Now().Add(time.Minute * time.Duration(3)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(APP_KEY))
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	cookie := http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Minute * 1),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/index", 302)
	return
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		token, error := jwt.Parse(auth.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("Error decrypting token")
				return nil, fmt.Errorf("Internal server error")
			}
			return []byte("secret"), nil
		})
		if error != nil {

			json.NewEncoder(w).Encode(error.Error())
			return
		}
		if token.Valid {
			next.ServeHTTP(w, r)
		}
	})
}
