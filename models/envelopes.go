package models

// HomeEnvelope :
type HomeEnvelope struct {
	User    User
	Quizzes []Quiz
	Offset  int
	Total   int
}

// ExecuteQuizzEnvelope :
type ExecuteQuizzEnvelope struct {
	User User
	Quiz Quiz
}
