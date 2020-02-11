package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Matias-Barrios/QuizApp/models"
	quizzes "github.com/Matias-Barrios/QuizApp/quizzes"
	"github.com/Matias-Barrios/QuizApp/views"
	"github.com/dgrijalva/jwt-go"
)

// Hanlders definitions
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/index" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	claims := getClaims(w, r)
	var offset int
	var offsetv int64
	var err error
	keys, ok := r.URL.Query()["offset"]
	if !ok || len(keys[0]) < 1 {
		offset = 0
	}
	if len(keys) > 0 {
		offsetv, err = strconv.ParseInt(keys[0], 10, 64)
	} else {
		offset = 0
	}

	if err != nil {
		offset = 0
	} else {
		offset = int(offsetv)
	}
	qs, err := quizzes.GetQuizzes(offset)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	u := models.User{
		Name: claims.Username,
	}
	envelope := models.HomeEnvelope{
		User:    u,
		Quizzes: qs,
	}
	views.ViewIndex.Execute(w, &envelope)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	views.ViewLogin.Execute(w, nil)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/favicon.ico" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "static/favicon.ico")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	notoken := http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(time.Minute * 800),
	}
	http.SetCookie(w, &notoken)
	http.Redirect(w, r, "/login", 302)
}

func executeQuizzHanlder(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/execute" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	claims := getClaims(w, r)
	u := models.User{
		Name: claims.Username,
	}
	keys, ok := r.URL.Query()["quizz"]
	if !ok || len(keys) < 1 || len(keys[0]) < 1 {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	quizz, err := quizzes.GetQuizzByID(keys[0])
	if err != nil {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	envelope := models.ExecuteQuizzEnvelope{
		User: u,
		Quiz: quizz,
	}

	views.ViewExecuteQuizz.Execute(w, &envelope)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		views.View404.Execute(w, nil)
	}
}

func getClaims(w http.ResponseWriter, r *http.Request) models.Claim {
	token, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return models.Claim{}
	}
	claims := &models.Claim{}
	_, err = jwt.ParseWithClaims(token.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return APP_KEY, nil
	})
	return *claims
}
