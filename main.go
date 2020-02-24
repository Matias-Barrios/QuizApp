package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Matias-Barrios/QuizApp/config"
	"github.com/Matias-Barrios/QuizApp/router"
)

func main() {
	f, err := os.OpenFile("/var/log/quizapp.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	envF := config.EnvironmentFetcher{}
	port, err := envF.GetValue("PORT")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Starting app in port : ", port)
	http.ListenAndServe(":"+port, router.GetRouter())
}
