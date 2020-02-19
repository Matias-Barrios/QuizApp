function ValidateUsername(textBox) {
    let register_button = document.getElementById('register');

    if (__validateUserName(textBox.value)){
        textBox.style.background = "#53bd66";
        register_button.disabled = false
    }
    else{
        textBox.style.background = "#d1736d";
        register_button.disabled = true;
    }
}

function ValidateEmail(textBox) {
    let register_button = document.getElementById('register');
    if (__validateEmail(textBox.value)){
        textBox.style.background = "#53bd66";
        delete register_button.disabled;
    }
    else{
        textBox.style.background = "#d1736d";
        register_button.disabled = true;
    }
}

function ValidatePassword(textBox) {
    let register_button = document.getElementById('register');
    let passwordRules ={
        password_validation_length :  document.getElementById('password_validation_length'),
        password_validation_oneupper :  document.getElementById('password_validation_oneupper'),
        password_validation_onelower :  document.getElementById('password_validation_onelower'),
        password_validation_special :  document.getElementById('password_validation_special')
    }
    if (__testPasswordisAtLeast8Characters(textBox.value)){
        passwordRules["password_validation_length"].style.color = "#53bd66";
        delete register_button.disabled;
    }
    else{
        passwordRules["password_validation_length"].style.color = "#d1736d";
        register_button.disabled = true;
    }
    
    if (__testPasswordisHasAtLeastOneLowerCaseCharacter(textBox.value)){ 
        passwordRules["password_validation_onelower"].style.color = "#53bd66";
        delete register_button.disabled;
    }
    else { 
        passwordRules["password_validation_onelower"].style.color = "#d1736d";
        register_button.disabled = true;
    }
    
    if (__testPasswordisHasAtLeastOneUpperCaseCharacter(textBox.value)){ 
        passwordRules["password_validation_oneupper"].style.color = "#53bd66";
        delete register_button.disabled;
    }
    else{ 
        passwordRules["password_validation_oneupper"].style.color = "#d1736d";
        register_button.disabled = true;
    }

    if (__testPasswordisHasAtLeastOneSpecialCharacter(textBox.value)){ 
        passwordRules["password_validation_special"].style.color = "#53bd66";
        delete register_button.disabled;
    }
    else {
        passwordRules["password_validation_special"].style.color = "#d1736d";
        register_button.disabled = true;
    }
    
    let passwordInput = document.getElementById('Pass');
    if (__testPasswordisAtLeast8Characters(textBox.value) &&
        __testPasswordisHasAtLeastOneLowerCaseCharacter(textBox.value) &&
        __testPasswordisHasAtLeastOneUpperCaseCharacter(textBox.value) &&
        __testPasswordisHasAtLeastOneSpecialCharacter(textBox.value)){ 
        passwordInput.style.background = "#53bd66";
        delete register_button.disabled;    
    }
    else{
        register_button.disabled = true
        passwordInput.style.background = "#d1736d";
    }
}

function ValidateRepeat(textBox){
    let register_button = document.getElementById('register');
    if (textBox == null) {
        register_button.disabled = true;
        return;
    }
    if ( __validateRepeat(textBox.value) ){
        delete register_button.disabled;
    }else{
        register_button.disabled = true;
    }
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


