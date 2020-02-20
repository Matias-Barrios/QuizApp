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
		SELECT Quizzes.id,
			   if(Users_Completed_Quizzes.user_id IS NOT NULL, true, false) as completed,
			   Quizzes.Content,
			   ( SELECT count(*) FROM Quizzes) as count 
			   FROM Quizzes LEFT JOIN Users_Completed_Quizzes 
			   ON Users_Completed_Quizzes.user_id = ?
			   AND Quizzes.id = Users_Completed_Quizzes.quiz_id
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

//  SetQuizzAsCompleted :
func SetQuizzAsCompleted(user_id int, id string) error {
	_, err := sqlConnection.Exec(`
		INSERT INTO Users_Completed_Quizzes
		(user_id,quiz_id)
		VALUES (?, ?)
		`, user_id, id)
	if err != nil {
		return err
	}
	return nil
}
