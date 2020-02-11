package models

// HomeEnvelope :
type HomeEnvelope struct {
	User    User
	Quizzes []Quiz
}

// ExecuteQuizzEnvelope :
type ExecuteQuizzEnvelope struct {
	User User
	Quiz Quiz
}
