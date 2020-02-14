let usernameInput = document.getElementById('Name');
let emailInput = document.getElementById('Email');
let passwordInput = document.getElementById('Pass');
let password_checkInput = document.getElementById('PassRepeat');


function ValidateUsername(textBox) {
    if (__validateUserName(textBox.value))
        textBox.style.background = "#53bd66";
    else
        textBox.style.background = "#d1736d";
}

function ValidateEmail(textBox) {
    if (__validateEmail(textBox.value))
        textBox.style.background = "#53bd66";
    else
        textBox.style.background = "#d1736d";
}

function ValidatePassword(textBox) {
    let passwordRules ={
        password_validation_length :  document.getElementById('password_validation_length'),
        password_validation_oneupper :  document.getElementById('password_validation_oneupper'),
        password_validation_onelower :  document.getElementById('password_validation_onelower'),
        password_validation_special :  document.getElementById('password_validation_special')
    }
    if (__testPasswordisAtLeast8Characters(textBox.value))
        passwordRules["password_validation_length"].style.color = "#53bd66";
    else
        passwordRules["password_validation_length"].style.color = "#d1736d";
    
    if (__testPasswordisHasAtLeastOneLowerCaseCharacter(textBox.value))
        passwordRules["password_validation_onelower"].style.color = "#53bd66";
    else
        passwordRules["password_validation_onelower"].style.color = "#d1736d";
    
    if (__testPasswordisHasAtLeastOneUpperCaseCharacter(textBox.value))
        passwordRules["password_validation_oneupper"].style.color = "#53bd66";
    else
        passwordRules["password_validation_oneupper"].style.color = "#d1736d";

    if (__testPasswordisHasAtLeastOneSpecialCharacter(textBox.value))
        passwordRules["password_validation_special"].style.color = "#53bd66";
    else
        passwordRules["password_validation_special"].style.color = "#d1736d";
    
    let passwordInput = document.getElementById('Pass');
    if (__testPasswordisAtLeast8Characters(textBox.value) &&
        __testPasswordisHasAtLeastOneLowerCaseCharacter(textBox.value) &&
        __testPasswordisHasAtLeastOneUpperCaseCharacter(textBox.value) &&
        __testPasswordisHasAtLeastOneSpecialCharacter(textBox.value))
        passwordInput.style.background = "#53bd66";
    else
        passwordInput.style.background = "#d1736d";
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
    console.log("Adentro : "+ input);
    var re = /^(?=.*[a-z]).*$/;
    console.log(re.test(String(input)))
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


