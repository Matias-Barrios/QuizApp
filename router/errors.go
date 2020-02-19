package router

import (
	"log"
	"net/http"

	"github.com/Matias-Barrios/QuizApp/views"
)

func internalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/error" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := views.ViewInternalServerError.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
	}
}
