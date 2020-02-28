let OKCOLOR = "#b5ffb8";
let BADCOLOR =  "#d1736d";
let TEXTOKCOLOR = "#49c94e";
let BUTTONOKCOLOR = "#49c94e";


function __changeButtonColors(color){
    document.getElementById('register').style.background = color;
}

function ValidateFields() {
    let register_button = document.getElementById('register');
    let textBox = document.getElementById('Pass');
    let repeatTextBox = document.getElementById('PassRepeat');
    let usernametextBox = document.getElementById('Name');
    let emailtextBox = document.getElementById('Email');

    if (__validateUserName(usernametextBox.value)){
        usernametextBox.style.background = OKCOLOR;
    }
    else{
        usernametextBox.style.background = BADCOLOR;
    }

    if (__validateEmail(emailtextBox.value)){
        emailtextBox.style.background = OKCOLOR;
    }
    else{
        emailtextBox.style.background = BADCOLOR;
    }
    let passwordRules ={
        password_validation_length :  document.getElementById('password_validation_length'),
        password_validation_oneupper :  document.getElementById('password_validation_oneupper'),
        password_validation_onelower :  document.getElementById('password_validation_onelower'),
        password_validation_special :  document.getElementById('password_validation_special')
    }
    if (__testPasswordisAtLeast8Characters(textBox.value)){
        passwordRules["password_validation_length"].style.color = TEXTOKCOLOR;
    }
    else{
        passwordRules["password_validation_length"].style.color = BADCOLOR;
    }
    
    if (__testPasswordisHasAtLeastOneLowerCaseCharacter(textBox.value)){ 
        passwordRules["password_validation_onelower"].style.color = TEXTOKCOLOR;
    }
    else { 
        passwordRules["password_validation_onelower"].style.color = BADCOLOR;
    }
    
    if (__testPasswordisHasAtLeastOneUpperCaseCharacter(textBox.value)){ 
        passwordRules["password_validation_oneupper"].style.color = TEXTOKCOLOR;
    }
    else{ 
        passwordRules["password_validation_oneupper"].style.color = BADCOLOR;
    }

    if (__testPasswordisHasAtLeastOneSpecialCharacter(textBox.value)){ 
        passwordRules["password_validation_special"].style.color = TEXTOKCOLOR;
    }
    else {
        passwordRules["password_validation_special"].style.color = BADCOLOR;
    }
    
    let passwordInput = document.getElementById('Pass');
    if (__testPasswordisAtLeast8Characters(textBox.value) &&
        __testPasswordisHasAtLeastOneLowerCaseCharacter(textBox.value) &&
        __testPasswordisHasAtLeastOneUpperCaseCharacter(textBox.value) &&
        __testPasswordisHasAtLeastOneSpecialCharacter(textBox.value)){ 
        passwordInput.style.background = OKCOLOR;
    }
    else{
        passwordInput.style.background = BADCOLOR;
    }
    if ( __validateRepeat(repeatTextBox.value)){
        repeatTextBox.style.background = OKCOLOR;
    }else{
        repeatTextBox.style.background = BADCOLOR;
    }

    if (__testPasswordisAtLeast8Characters(textBox.value) &&
        __testPasswordisHasAtLeastOneLowerCaseCharacter(textBox.value) &&
        __testPasswordisHasAtLeastOneUpperCaseCharacter(textBox.value) &&
        __testPasswordisHasAtLeastOneSpecialCharacter(textBox.value) &&
        __validateEmail(emailtextBox.value) &&
        __validateUserName(usernametextBox.value) &&
        __validateRepeat(repeatTextBox.value)){
            register_button.disabled = false
            __changeButtonColors(BUTTONOKCOLOR);
        }
    else{
            register_button.disabled = true
            __changeButtonColors(BADCOLOR);
    }
}


function Submit(){
    let password = document.getElementById('Pass').value;
    let usernametext = document.getElementById('Name').value;
    let emailtext = document.getElementById('Email').value;
    let captchaid = document.getElementById('captchadiv').getAttribute('value');
    let captchasolution = document.getElementById('CaptchaSolution').value;
    
    if (password == null || usernametext == null || emailtext == null)
        return;
    fetch('/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username : usernametext,
            email : emailtext,
            password : password,
            captchaid : parseInt(captchaid),
            solution : captchasolution
        })
    })
    .then(function(response) {
        if (response.status != 200) {
            window.location.replace("/error");
        }else {
        window.location.replace("/success");
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

function __validateUserName(input) {
    var re = /^[a-zA-Z][a-zA-Z0-9_-]{5,}$/;
    return re.test(String(input));
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


