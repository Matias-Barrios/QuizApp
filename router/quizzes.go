package router

import (
	"encoding/json"
	"net/http"

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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(solution)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solution)
}
