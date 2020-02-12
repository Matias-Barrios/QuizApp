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
    console.log(answer);
}

document.addEventListener("DOMContentLoaded", function(event) { 
    // TODO
});