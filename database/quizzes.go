package database

import (
	"encoding/json"
	"strconv"

	"github.com/Matias-Barrios/QuizApp/models"
)

func GetQuizzes(userid, offset int) ([]models.Quiz, int, error) {
	var quizzes []models.Quiz
	quizzes = make([]models.Quiz, 0)
	rows, err := sqlConnection.Query(`
		SELECT id,if(Users_Completed_Quizzes.user_id is not null, true, false) as completed, content, (
			SELECT count(*) FROM Quizzes
		) as count
		FROM Quizzes
		LEFT OUTER JOIN Users_Completed_Quizzes ON 
		Users_Completed_Quizzes.user_id = ? AND  Users_Completed_Quizzes.user_id = id
		AND Quizzes.active = true
		LIMIT 10 OFFSET ?
		`, userid, offset)
	if err != nil {
		return nil, 0, err
	}
	var count int
	for rows.Next() {
		var id int
		var content string
		var completed bool
		err := rows.Scan(&id, &completed, &content, &count)
		if err != nil {
			return nil, 0, err
		}
		var q models.Quiz
		err = json.Unmarshal([]byte(content), &q)
		if err != nil {
			return nil, 0, err
		}
		q.ID = strconv.Itoa(id)
		q.Completed = completed
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
