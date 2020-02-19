package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Matias-Barrios/QuizApp/config"
)

var sqlConnection *sql.db

func init() {
	sqlConnection = CreateConnection()
}

// CreateConnection :
func CreateConnection() *sql.db {
	envF := config.EnvironmentFetcher{}
	dbuser, err := envF.GetValue("DBUSER")
	if err != nil {
		log.Fatalln(err.Error())
	}
	dbpassword, err := envF.GetValue("DBPASSWORD")
	if err != nil {
		log.Fatalln(err.Error())
	}
	dbaddress, err := envF.GetValue("DBADDRESS")
	if err != nil {
		log.Fatalln(err.Error())
	}
	dbport, err := envF.GetValue("DBPORT")
	if err != nil {
		log.Fatalln(err.Error())
	}
	dbnamespace, err := envF.GetValue("DBNAMESPACE")
	if err != nil {
		log.Fatalln(err.Error())
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpassword, dbaddress, dbport, dbnamespace))
	if err != nil {
		panic(err.Error())
	}
	return db
}
