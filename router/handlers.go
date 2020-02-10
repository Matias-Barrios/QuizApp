package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Matias-Barrios/QuizApp/models"
	quizzes "github.com/Matias-Barrios/QuizApp/quizzes"
	homeV "github.com/Matias-Barrios/QuizApp/views"
	loginV "github.com/Matias-Barrios/QuizApp/views"
)

// Hanlders definitions
func indexHandler(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	var offset int
	var offsetv int64
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
		Name: username.Value,
	}
	envelope := models.HomeEnvelope{
		User:    u,
		Quizzes: qs,
	}
	homeV.ViewIndex.Execute(w, &envelope)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	loginV.ViewLogin.Execute(w, nil)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon.ico")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	notoken := http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(time.Minute * 800),
	}
	http.SetCookie(w, &notoken)
	http.Redirect(w, r, "/login", 302)
}
