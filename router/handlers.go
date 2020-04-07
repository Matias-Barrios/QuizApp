package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Matias-Barrios/QuizApp/database"
	"github.com/Matias-Barrios/QuizApp/models"
	"github.com/Matias-Barrios/QuizApp/views"
)

// Hanlders definitions
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/index" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	claims := getClaims(w, r)
	var offset int
	var offsetv int64
	var err error
	keys, ok := r.URL.Query()["offset"]
	if !ok || len(keys[0]) < 1 {
		offset = 0
	}
	if len(keys) > 0 {
		offsetv, err = strconv.ParseInt(keys[0], 10, 64)
	} else {
		offset = 0
	}

	if err != nil {
		log.Println(err.Error())
		offset = 0
	} else {
		offset = int(offsetv)
	}
	qs, count, err := database.GetQuizzes(claims.User.ID, offset)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/login", 302)
		return
	}
	u := models.User{
		Name: claims.User.Name,
	}
	envelope := models.HomeEnvelope{
		User:    u,
		Offset:  offset,
		Total:   count,
		Quizzes: qs,
	}
	err = views.ViewIndex.Execute(w, &envelope)
	if err != nil {
		log.Println(err.Error())
	}
	err = database.Log(r.RemoteAddr, claims.User.Email, time.Now().UTC().Unix(), "VISITED", "User has landed on the home page")
	if err != nil {
		log.Println(err.Error())
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	lerror := models.LoginError{}
	keys, ok := r.URL.Query()["error"]
	if ok && len(keys[0]) > 0 {
		lerror.Message = MapErrorMessage(keys[0])
	}
	err := views.ViewLogin.Execute(w, lerror)
	if err != nil {
		log.Println(err.Error())
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/favicon.ico" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "static/favicon.ico")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	notoken := http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(time.Minute * 800),
	}
	err := database.Log(r.RemoteAddr, "", time.Now().UTC().Unix(), "LOGOUT", "A user has logged out")
	if err != nil {
		log.Println(err.Error())
	}
	http.SetCookie(w, &notoken)
	http.Redirect(w, r, "/login", 302)
}

func executeQuizzHanlder(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/execute" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	claims := getClaims(w, r)
	u := models.User{
		Name: claims.User.Name,
	}
	keys, ok := r.URL.Query()["quizz"]
	if !ok || len(keys) < 1 || len(keys[0]) < 1 {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	quizz, err := database.GetQuizzByID(keys[0])
	if err != nil {
		log.Println(err.Error())
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	quizz.ID = keys[0]
	envelope := models.ExecuteQuizzEnvelope{
		User: u,
		Quiz: quizz,
	}

	err = views.ViewExecuteQuizz.Execute(w, &envelope)

	if err != nil {
		log.Println(err.Error())
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		views.View404.Execute(w, nil)
	}
	if status == http.StatusInternalServerError {
		views.ViewInternalServerError.Execute(w, nil)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	id, captchapath, err := database.GenerateCapctha(r.RemoteAddr)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/error", 302)
		return
	}
	captcha := models.RegisterCaptcha{
		ID:   id,
		Path: captchapath,
	}
	previousFailure, ok := r.URL.Query()["error"]
	if ok && len(previousFailure) > 0 {
		captcha.PreviousFailure = previousFailure[0]
	}
	err = views.ViewRegister.Execute(w, captcha)
	if err != nil {
		log.Println(err.Error())
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/create" && r.Method != "POST" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	registerBody := models.RegisterBody{}
	err := decoder.Decode(&registerBody)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		http.Redirect(w, r, "/error", 302)
		return
	}
	validusername := regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_-]{5,}$")
	validemail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	eightormore, lower, upper, symbol := verifyPassword(registerBody.Password)
	if !eightormore || !lower || !upper || !symbol || !validusername.Match([]byte(registerBody.Username)) || !validemail.Match([]byte(registerBody.Email)) {
		w.WriteHeader(http.StatusBadRequest)
		http.Redirect(w, r, "/error", 302)
		return
	}
	captcha, err := database.GetCaptcha(registerBody.CaptchaID)
	if err != nil || strings.ToLower(strings.TrimSpace(registerBody.Solution)) != strings.ToLower(strings.TrimSpace(captcha.Result)) {
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println(fmt.Sprintf("Failed to test captcha from %s with result %s", r.RemoteAddr, registerBody.Solution))
		}
		http.Redirect(w, r, "/register?error=2", 302)
		return
	}
	err = database.CreateUser(registerBody.Username, registerBody.Password, registerBody.Email)
	if err != nil {
		serr := database.Log(r.RemoteAddr, registerBody.Email, time.Now().UTC().Unix(), "USERCREATIONERROR", err.Error())
		if serr != nil {
			log.Println(serr.Error())
		}
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		err = database.Log(r.RemoteAddr, registerBody.Email, time.Now().UTC().Unix(), "USERCREATED", "A new user has been created")
		if err != nil {
			log.Println(err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{ "status" : "success" }`))
		return
	}
}

func successCreationHanlder(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/success" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := views.ViewSuccessCreation.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
	}
}

func changePasswordHanlder(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/changepassword" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := views.ViewChangePassword.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := views.ViewAbout.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
	}
}
