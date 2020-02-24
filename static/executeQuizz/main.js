function validate(ID) {
    var divs  = document.getElementsByClassName('question_container');
    var answer = {
        quizID : ID,
        answers : []
    }
    for (let item of divs) {
        questionID = [].slice.call(item.getElementsByTagName('input'))[0]?.name;     
        checked = [].slice.call(item.getElementsByTagName('input'))
                    .filter(element => element.checked)
                    .map(item => item.value);
        answer.answers.push({
            questionID : questionID,
            values : checked
        });
    }
    (async () => {
        const rawResponse = await fetch('/validate', {
          method: 'POST',
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(answer)
        });
        console.log("Coloring!");
        const solved = await rawResponse.json();
        solved.answers.forEach(element => {
          if (element.passed) {
            document.getElementById('question_' + element.questionID).style.background="#4bc96c";
          }else{
            document.getElementById('question_' + element.questionID).style.background="#e86d6d";
          }
        });

        document.getElementById('validate_quiz').style.display = "none";
        let percentageBanner = document.getElementById('percentage_completed');
        console.log(percentageBanner)
        if (percentageBanner != null) { 
          percentageBanner.style.display = "inline-block";
          if (solved.percentageCompleted >= 100)
            percentageBanner.style.background = "#4bc96c"; 
          else 
            percentageBanner.style.background = "#d9d779"; 
          percentageBanner
            .querySelector("#percentage_completed_text")
            .textContent = `Your score for this test : ${solved.percentageCompleted}`;
        }
      })();
}

document.addEventListener("DOMContentLoaded", function(event) { 
    // TODO
});