package router

import (
	"io"
	"log"
	"net/http"
	"time"

	homeV "github.com/Matias/QuizApp/views/view_index"
	loginV "github.com/Matias/QuizApp/views/view_login"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

const (
	APP_KEY = "golangcode.com"
)

// GetRouter :
func GetRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	mux.Handle("/index", AuthMiddleware(http.HandlerFunc(loginHandler)))
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
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"error":"invalid_credentials"}`)
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
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"token_generation_failed"}`)
		return
	}
	io.WriteString(w, `{"token":"`+tokenString+`"}`)
	return
}

func AuthMiddleware(next http.Handler) http.Handler {
	if len(APP_KEY) == 0 {
		log.Fatal("HTTP server unable to start, expected an APP_KEY for JWT auth")
	}
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(APP_KEY), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return jwtMiddleware.Handler(next)
}
