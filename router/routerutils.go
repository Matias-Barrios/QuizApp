package router

import (
	"unicode"

	"github.com/Matias-Barrios/QuizApp/database"
	"github.com/Matias-Barrios/QuizApp/models"
)

func validate(user_id int, quiz models.Quiz, solution *models.Solution) error {
	for ix := range quiz.Questions {
		for subix := range solution.Answers {
			if quiz.Questions[ix].ID == solution.Answers[subix].QuestionID {
				if compareStringSlices(quiz.Questions[ix].Answers, solution.Answers[subix].Values) {
					solution.Answers[subix].Passed = true
				} else {
					solution.Answers[subix].Passed = false
				}
			}
		}
	}
	solution.PercentageCompleted = averageCompleted(*solution)
	if solution.PercentageCompleted > 80 {
		err := database.SetQuizzAsCompleted(user_id, quiz.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func averageCompleted(solution models.Solution) int {
	var count = countCompleted(solution)
	if count == 0 {
		return 0
	}
	return count * 100 / len(solution.Answers)
}

func countCompleted(solution models.Solution) int {
	var counter int
	for _, s := range solution.Answers {
		if s.Passed {
			counter++
		}
	}
	return counter
}

func compareStringSlices(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var counter int
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				counter++
			}
		}
	}
	if counter == len(s1) {
		return true
	}
	return false
}

func verifyPassword(s string) (eightOrMore, lower, upper, special bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsLower(c):
			lower = true
			letters++
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
			letters++
		case unicode.IsLetter(c) || c == ' ':
			letters++
		case unicode.IsNumber(c):
			letters++
		default:
			return false, false, false, false
		}
	}
	eightOrMore = letters >= 8
	return
}
