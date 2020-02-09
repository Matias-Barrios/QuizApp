package router

import (
	"net/http"

	homeV "github.com/Matias-Barrios/QuizApp/views"
	loginV "github.com/Matias-Barrios/QuizApp/views"
)

// Hanlders definitions
func indexHandler(w http.ResponseWriter, r *http.Request) {
	homeV.ViewIndex.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	loginV.ViewLogin.Execute(w, nil)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon.ico")
}
