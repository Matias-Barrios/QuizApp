

SELECT id,if(Users_Completed_Quizzes.user_id is not null, true, false) as completed, (
	SELECT count(*) FROM Quizzes
) as count
FROM Quizzes
LEFT OUTER JOIN Users_Completed_Quizzes ON Users_Completed_Quizzes.user_id = id 
AND Quizzes.active = true;


INSERT INTO Users_Completed_Quizzes VALUES(1,1);