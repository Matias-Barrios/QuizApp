package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Matias-Barrios/QuizApp/database"
	"github.com/Matias-Barrios/QuizApp/models"
	"github.com/Matias-Barrios/QuizApp/views"
	"github.com/dgrijalva/jwt-go"
)

// Claims :
var Claims models.Claim

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := database.GetUser(password, username)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/login", 302)
		return
	}
	claims := &models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(800 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(APP_KEY))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/login", 302)
		return
	}
	cookie := http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Minute * 800),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/index", 302)
	return
}

func getClaims(w http.ResponseWriter, r *http.Request) models.Claim {
	token, err := r.Cookie("token")
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/login", 302)
		return models.Claim{}
	}
	claims := &models.Claim{}
	_, err = jwt.ParseWithClaims(token.Value, claims, func(token *jwt.Token) (interface{}, error) {
		//log.Println(err.Error())
		return APP_KEY, nil
	})
	return *claims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, err := r.Cookie("token")
		if err != nil {
			log.Println(err.Error())
			http.Redirect(w, r, "/login", 302)
			return
		}
		token, err := jwt.Parse(auth.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("Error decrypting token")
				return nil, fmt.Errorf("Internal server error")
			}
			return []byte(APP_KEY), nil
		})
		if err != nil {
			log.Println(err.Error())
			http.Redirect(w, r, "/login", 302)
			return
		}
		if token.Valid {
			next.ServeHTTP(w, r)
		}
	})
}

func forgotHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/forgot" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := views.ViewForgot.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
	}
}
