package router

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Matias-Barrios/QuizApp/database"
	"github.com/Matias-Barrios/QuizApp/models"
)

func validateQuizzHanlder(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/validate" && r.Method != "POST" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	claims := getClaims(w, r)
	decoder := json.NewDecoder(r.Body)
	var solution models.Solution
	err := decoder.Decode(&solution)
	if err != nil {
		errorHandler(w, r, http.StatusNotFound)
	}
	quizz, err := database.GetQuizzByID(solution.QuizID)
	if err != nil {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err = validate(claims.User.ID, quizz, &solution)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(solution)
		return
	}
	if solution.PercentageCompleted > 80 {
		err = database.Log(r.RemoteAddr, claims.User.Email, time.Now().UTC().Unix(), "QUIZCOMPLETED", "User has completed a Quiz.")
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		err = database.Log(r.RemoteAddr, claims.User.Email, time.Now().UTC().Unix(), "QUIZFAILED", "User has failed to complete a Quiz.")
		if err != nil {
			log.Println(err.Error())
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solution)
}
