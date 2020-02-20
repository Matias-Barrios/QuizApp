package database

import (
	"encoding/json"
	"strconv"

	"github.com/Matias-Barrios/QuizApp/models"
)

func GetQuizzes(offset int) ([]models.Quiz, int, error) {
	var quizzes []models.Quiz
	quizzes = make([]models.Quiz, 0)
	rows, err := sqlConnection.Query(`
		SELECT id, content, (
			SELECT count(*) FROM Quizzes
		) as count
		FROM Quizzes
		WHERE active = true
		LIMIT 10 OFFSET ?
		`, offset)
	if err != nil {
		return nil, 0, err
	}
	var count int
	for rows.Next() {
		var id int
		var content string
		err := rows.Scan(&id, &content, &count)
		if err != nil {
			return nil, 0, err
		}
		var q models.Quiz
		err = json.Unmarshal([]byte(content), &q)
		if err != nil {
			return nil, 0, err
		}
		q.ID = strconv.Itoa(id)
		quizzes = append(quizzes, q)
	}
	return quizzes, count, nil
}

// GetQuizzByID :
func GetQuizzByID(id string) (models.Quiz, error) {
	var data string
	err := sqlConnection.QueryRow(`
		SELECT content
		FROM Quizzes
		WHERE active = true
		AND ID = ?
		`, id).Scan(&data)

	var quizz models.Quiz
	err = json.Unmarshal([]byte(data), &quizz)
	if err != nil {
		return models.Quiz{}, err
	}
	quizz.ID = id
	return quizz, err
}
