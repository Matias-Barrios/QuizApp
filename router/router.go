package router

import (
	"log"
	"net/http"

	"github.com/Matias-Barrios/QuizApp/config"
)

var APP_KEY string

func init() {
	var err error
	envF := config.EnvironmentFetcher{}
	APP_KEY, err = envF.GetValue("APP_KEY")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// GetRouter :
func GetRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	mux.Handle("/index", AuthMiddleware(http.HandlerFunc(indexHandler)))
	mux.HandleFunc("/auth", TokenHandler)
	mux.HandleFunc("/favicon.ico", faviconHandler)
	mux.HandleFunc("/logout", logoutHandler)

	// Static files handling
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return mux
}
