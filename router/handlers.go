package router

import (
	"net/http"
	"time"

	"github.com/Matias-Barrios/QuizApp/models"
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
	u := &models.User{
		Name: username.Value,
	}
	homeV.ViewIndex.Execute(w, u)
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
