DROP DATABASE IF EXISTS QUIZAPP;
CREATE DATABASE IF NOT EXISTS QUIZAPP;

use QUIZAPP;

CREATE TABLE Users (
     id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
     name VARCHAR(50) NOT NULL,
     email VARCHAR(50) NOT NULL,
     password_encrypted VARCHAR(150) NOT NULL,
     created_on BIGINT UNSIGNED NOT NULL,
     active BOOL NOT NULL,
     CONSTRAINT unique_email UNIQUE (email)
);

CREATE TABLE Quizzes (
     id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
     content TEXT NOT NULL,
     active BOOL NOT NULL,
     CHECK (JSON_VALID(content))
);

CREATE TABLE Users_Completed_Quizzes (
  user_id INT UNSIGNED NOT NULL,
  quiz_id INT UNSIGNED NOT NULL,
  PRIMARY KEY (user_id,quiz_id),
  FOREIGN KEY (user_id) REFERENCES Users (id),
  FOREIGN KEY (quiz_id) REFERENCES Quizzes (id)
);


CREATE TABLE Unsuccessful_login_attempts (
  user_id INT UNSIGNED  NOT NULL,
  attempted_on BIGINT UNSIGNED NOT NULL,
  PRIMARY KEY (user_id,attempted_on),
  FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE TABLE LOGS (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  remote_ip VARCHAR(50) NOT NULL,
  email VARCHAR(50) NOT NULL,
  occurred_on BIGINT UNSIGNED NOT NULL,
  comment VARCHAR(120) NOT NULL
);


DELIMITER $$
CREATE TRIGGER check_username_validity BEFORE INSERT ON Users
FOR EACH ROW 
BEGIN 
IF (NEW.name REGEXP '^[a-zA-Z0-9_-]{5,20}$' ) = 0 THEN 
  SIGNAL SQLSTATE '12345'
     SET MESSAGE_TEXT = 'Wrong username format!';
END IF; 
END$$
DELIMITER ;
