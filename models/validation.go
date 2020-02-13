package models

// Solution :
type Solution struct {
	QuizID  string    `json:"quizID"`
	Answers []Answers `json:"answers"`
}

// Answers :
type Answers struct {
	QuestionID string   `json:"questionID"`
	Values     []string `json:"values"`
	Passed     bool     `json:"passed"`
}
