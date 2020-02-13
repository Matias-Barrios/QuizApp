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
	solution.PercentageCompleted = averageCompleted(*solution)
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
