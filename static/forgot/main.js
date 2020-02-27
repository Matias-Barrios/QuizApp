let OKCOLOR = "#b5ffb8";
let BADCOLOR =  "#d1736d";
let TEXTOKCOLOR = "#49c94e";
let BUTTONOKCOLOR = "#49c94e";

function ValidateFields() {
    let emailtextBox = document.getElementById('Email');
    let sendemailButton = document.getElementById('sendemail');
    if (__validateEmail(emailtextBox.value)){
        emailtextBox.style.background = OKCOLOR;
    }
    else{
        emailtextBox.style.background = BADCOLOR;
    }
    if (__validateEmail(emailtextBox.value)){
            sendemailButton.disabled = false
            __changeButtonColors(BUTTONOKCOLOR);
        }
    else{
            sendemailButton.disabled = true
            __changeButtonColors(BADCOLOR);
    }
}


function Submit(){
    let sendEmailValue = document.getElementById('Email').value;
    if (sendEmailValue == null)
        return;
    fetch('/forgot', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            email : sendEmailValue })
    })
    .then(function(response) {
        if (response.status != 200) {
            window.location.replace("/error");
        }else {
            window.location.replace("/login");
        }
    })
    .then(function(_) {
    })
    .catch(function(err) {
        console.log("Redirecting to error "+ err)
        window.location.replace("/error");
    });
        
}

function __validateRepeat(input){
    let password_text = document.getElementById('Pass').value
    let repeat_text = input
    return (password_text != null && repeat_text != null ) && ( password_text == repeat_text )
}

function __validateEmail(input) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(input).toLowerCase());
}

function __testPasswordisAtLeast8Characters(input){
    var re = /^.{8,}$/;
    return re.test(String(input));
}

function __testPasswordisHasAtLeastOneLowerCaseCharacter(input){
    var re = /^(?=.*[a-z]).*$/;
    return re.test(String(input));
}

function __testPasswordisHasAtLeastOneUpperCaseCharacter(input){
    var re = /^(?=.*[A-Z]).*$/;
    return re.test(String(input));
}

function __testPasswordisHasAtLeastOneSpecialCharacter(input){
    var re = /^(?=.*[!@#$%^&*]).*$/;
    return re.test(String(input));
}

function __changeButtonColors(color){
    document.getElementById('sendemail').style.background = color;
}

