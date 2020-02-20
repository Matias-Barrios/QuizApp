package router

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/Matias-Barrios/QuizApp/database"
	"github.com/Matias-Barrios/QuizApp/models"
	quizzes "github.com/Matias-Barrios/QuizApp/quizzes"
	"github.com/Matias-Barrios/QuizApp/views"
	"github.com/dgrijalva/jwt-go"
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
		offset = 0
	} else {
		offset = int(offsetv)
	}
	qs, err := quizzes.GetQuizzes(offset)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	u := models.User{
		Name: claims.User.Name,
	}
	envelope := models.HomeEnvelope{
		User:    u,
		Quizzes: qs,
	}
	err = views.ViewIndex.Execute(w, &envelope)
	if err != nil {
		log.Println(err.Error())
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := views.ViewLogin.Execute(w, nil)
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
	quizz, err := quizzes.GetQuizzByID(keys[0])
	if err != nil {
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

func validateQuizzHanlder(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/validate" && r.Method != "POST" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var solution models.Solution
	err := decoder.Decode(&solution)
	if err != nil {
		errorHandler(w, r, http.StatusNotFound)
	}
	quizz, err := quizzes.GetQuizzByID(solution.QuizID)
	if err != nil {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	validate(quizz, &solution)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solution)
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

func getClaims(w http.ResponseWriter, r *http.Request) models.Claim {
	token, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return models.Claim{}
	}
	claims := &models.Claim{}
	_, err = jwt.ParseWithClaims(token.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return APP_KEY, nil
	})
	return *claims
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" && r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := views.ViewRegister.Execute(w, nil)
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
		http.Redirect(w, r, "/error", 302)
		return
	}
	validusername := regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_-]{5,}$")
	validemail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	eightormore, lower, upper, symbol := verifyPassword(registerBody.Password)
	if !eightormore || !lower || !upper || !symbol || !validusername.Match([]byte(registerBody.Username)) || !validemail.Match([]byte(registerBody.Email)) {
		http.Redirect(w, r, "/error", 302)
		return
	}
	err = database.CreateUser(registerBody.Username, registerBody.Password, registerBody.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
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
