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
	mux.HandleFunc("/login", loginHandler)
	mux.Handle("/index", AuthMiddleware(http.HandlerFunc(indexHandler)))
	mux.HandleFunc("/auth", TokenHandler)
	mux.HandleFunc("/favicon.ico", faviconHandler)
	mux.HandleFunc("/logout", logoutHandler)

	// Static files handling
	mux.Handle("/static/", fileServerWithCustom404(http.Dir("static")))
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
