package router

import (
	"log"
	"net/http"
	"os"
	"path"

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

	// Protected endpoints
	mux.Handle("/index", AuthMiddleware(http.HandlerFunc(indexHandler)))
	mux.Handle("/execute", AuthMiddleware(http.HandlerFunc(executeQuizzHanlder)))
	mux.Handle("/validate", AuthMiddleware(http.HandlerFunc(validateQuizzHanlder)))

	// Unprotected endpoints
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/auth", TokenHandler)
	mux.HandleFunc("/register", registerHandler)
	mux.HandleFunc("/favicon.ico", faviconHandler)
	mux.HandleFunc("/logout", logoutHandler)
	mux.HandleFunc("/create", createUserHandler)
	mux.HandleFunc("/error", internalServerErrorHandler)

	// Static files handling
	//mux.Handle("/static/", fileServerWithCustom404(http.Dir("static")))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return mux
}

func fileServerWithCustom404(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		fsh.ServeHTTP(w, r)
	})
}
