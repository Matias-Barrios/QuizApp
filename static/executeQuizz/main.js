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
        const content = await rawResponse.json();
      
        console.log(content);
      })();
}

document.addEventListener("DOMContentLoaded", function(event) { 
    // TODO
});