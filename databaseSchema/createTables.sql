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
     content TEXT NOT NULL
);

CREATE TABLE Users_Completed_Quizzes (
  user_id INT UNSIGNED NOT NULL,
  quiz_id INT UNSIGNED NOT NULL,
  PRIMARY KEY (user_id,quiz_id),
  FOREIGN KEY (user_id) REFERENCES Users (id),
  FOREIGN KEY (quiz_id) REFERENCES Quizzes (id)
);


DELIMITER $$
CREATE TRIGGER check_username_validity BEFORE INSERT ON Users
FOR EACH ROW 
BEGIN 
IF (NEW.name REGEXP '^[a-zA-Z0-9_-]{5,20}$' ) = 0 THEN 
  SIGNAL SQLSTATE '12345'
     SET MESSAGE_TEXT = 'Wroooong!!!';
END IF; 
END$$
DELIMITER ;


DELIMITER $$
CREATE TRIGGER check_quiz_content BEFORE INSERT ON Quizzes
FOR EACH ROW 
BEGIN 
IF (NEW.content REGEXP '^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$
' ) = 0 THEN 
  SIGNAL SQLSTATE '12345'
     SET MESSAGE_TEXT = 'Wroooong!!!';
END IF; 
END$$
DELIMITER ;