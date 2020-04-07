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
	qcertificatekey, err := envF.GetValue("QCERTIFICATEKEY")
	if err != nil {
		log.Fatalln(err.Error())
	}
	qcertificatecrt, err := envF.GetValue("QCERTIFICATECRT")
	if err != nil {
		log.Fatalln(err.Error())
	}
	port, err := envF.GetValue("PORT")
	if err != nil {
		log.Fatalln(err.Error())
	}
	go redirectToHTTPS()
	log.Println("Starting app in port : ", port)
	err = http.ListenAndServeTLS(":"+port, qcertificatecrt, qcertificatekey, router.GetRouter())
	log.Println(err.Error())
}

func redirectToHTTPS() {
	httpSrv := http.Server{
		Addr: ":80",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://linuxquizapp.com.uy"+r.URL.String(), http.StatusMovedPermanently)
		}),
	}
	log.Println(httpSrv.ListenAndServe())
}
