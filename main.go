package main

import (
	"log"
	"net/http"

	"github.com/Matias-Barrios/QuizApp/config"
	"github.com/Matias-Barrios/QuizApp/router"
)

func main() {
	envF := config.EnvironmentFetcher{}
	port, err := envF.GetValue("PORT")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Starting app in port : ", port)
	http.ListenAndServe(":"+port, router.GetRouter())
}
