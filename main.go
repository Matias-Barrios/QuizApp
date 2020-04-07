package main

import (
	"log"
	"net"
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
	_, tlsPort, err := net.SplitHostPort(":" + port)
	if err != nil {
		log.Println(err.Error())
		os.Exit(2)
	}
	go redirectToHTTPS(tlsPort)
	log.Println("Starting app in port : ", port)
	err = http.ListenAndServeTLS(":"+port, qcertificatecrt, qcertificatekey, router.GetRouter())
	log.Println(err.Error())
}

func redirectToHTTPS(tlsPort string) {
	httpSrv := http.Server{
		Addr: ":80",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			host, _, _ := net.SplitHostPort(r.Host)
			u := r.URL
			u.Host = net.JoinHostPort(host, tlsPort)
			u.Scheme = "https"
			log.Println(u.String())
			http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
		}),
	}
	log.Println(httpSrv.ListenAndServe())
}
