package database

import "log"

// id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
//   remote_ip VARCHAR(50) NOT NULL,
//   email VARCHAR(50) NOT NULL,
//   occurred_on BIGINT UNSIGNED NOT NULL,
//   comment VARCHAR(120) NOT NULL

// Log :
func Log(remoteIP string, email string, occurredOn int64, occurrenceType string, comment string) error {
	_, err := sqlConnection.Exec(`
		INSERT INTO LOGS 
		(remote_ip, email, occurred_on,occurrence_type,comment)
		VALUES(?, ?, ?, ?, ?)	
	`, remoteIP, email, occurredOn, occurrenceType, comment)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
