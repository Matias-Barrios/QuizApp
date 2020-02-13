package router

import (
	"github.com/Matias-Barrios/QuizApp/models"
)

func validate(quiz models.Quiz, solution *models.Solution) {
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
