package quizzes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/Matias-Barrios/QuizApp/config"
	"github.com/Matias-Barrios/QuizApp/models"
)

var directory string

func init() {
	var err error
	envF := config.EnvironmentFetcher{}
	directory, err = envF.GetValue("QUIZZESDIR")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// GetQuizzes :
func GetQuizzes(offset int) ([]models.Quiz, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(files) == 0 {
		return nil, nil
	}
	var quizzes []models.Quiz
	var slice []os.FileInfo
	var count = 0
	for _, item := range files {
		if count > 10 {
			break
		} else {
			count++
		}
		slice = append(slice, item)
	}

	if len(slice) > 0 {
		for _, f := range slice {
			var q models.Quiz
			fd, err := os.Open(directory + "/" + f.Name())
			if err != nil {
				return nil, err
			}
			bytes, err := ioutil.ReadAll(fd)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(bytes, &q)
			if err != nil {
				return nil, err
			}
			var extension = filepath.Ext(f.Name())
			q.ID = (f.Name())[0 : len(f.Name())-len(extension)]
			quizzes = append(quizzes, q)
		}
	}
	return quizzes, nil
}
