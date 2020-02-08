package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	homeV "github.com/Matias-Barrios/QuizApp/views/view_index"
	loginV "github.com/Matias-Barrios/QuizApp/views/view_login"
	"github.com/dgrijalva/jwt-go"
)

const (
	APP_KEY = "secret"
)

// GetRouter :
func GetRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	mux.Handle("/index", AuthMiddleware(http.HandlerFunc(indexHandler)))
	mux.HandleFunc("/auth", TokenHandler)
	return mux
}

// Hanlders definitions
func indexHandler(w http.ResponseWriter, r *http.Request) {
	homeV.ViewIndex.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	loginV.ViewLogin.Execute(w, nil)
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if username != "myusername" || password != "mypassword" {
		http.Redirect(w, r, "/login", 302)
		return
	}

	// We are happy with the credentials, so build a token. We've given it
	// an expiry of 1 hour.
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
		Name:    "jwt",
		Value:   tokenString,
		Expires: time.Now().AddDate(0, 0, 1),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/index", 302)
	return
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, err := r.Cookie("jwt")
		if err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		token, error := jwt.Parse(auth.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Internal server error")
			}
			return []byte("secret"), nil
		})
		if error != nil {
			json.NewEncoder(w).Encode("Internal server error")
			return
		}
		if token.Valid {
			next.ServeHTTP(w, r)
		}
	})
}
