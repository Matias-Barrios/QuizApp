let OKCOLOR = "#b5ffb8";
let BADCOLOR =  "#d1736d";
let TEXTOKCOLOR = "#49c94e";
let BUTTONOKCOLOR = "#49c94e";


function __changeButtonColors(color){
    document.getElementById('change').style.background = color;
}

function ValidateFields() {
    let change_button = document.getElementById('change');
    let newpassword = document.getElementById('Pass');
    let repeatnewpassword = document.getElementById('PassRepeat');
    let currentpassword = document.getElementById('Current');

    let passwordRules ={
        password_validation_length :  document.getElementById('password_validation_length'),
        password_validation_oneupper :  document.getElementById('password_validation_oneupper'),
        password_validation_onelower :  document.getElementById('password_validation_onelower'),
        password_validation_special :  document.getElementById('password_validation_special')
    }
    if (__testPasswordisAtLeast8Characters(newpassword.value)){
        passwordRules["password_validation_length"].style.color = TEXTOKCOLOR;
    }
    else{
        passwordRules["password_validation_length"].style.color = BADCOLOR;
    }
    
    if (__testPasswordisHasAtLeastOneLowerCaseCharacter(newpassword.value)){ 
        passwordRules["password_validation_onelower"].style.color = TEXTOKCOLOR;
    }
    else { 
        passwordRules["password_validation_onelower"].style.color = BADCOLOR;
    }
    
    if (__testPasswordisHasAtLeastOneUpperCaseCharacter(newpassword.value)){ 
        passwordRules["password_validation_oneupper"].style.color = TEXTOKCOLOR;
    }
    else{ 
        passwordRules["password_validation_oneupper"].style.color = BADCOLOR;
    }

    if (__testPasswordisHasAtLeastOneSpecialCharacter(newpassword.value)){ 
        passwordRules["password_validation_special"].style.color = TEXTOKCOLOR;
    }
    else {
        passwordRules["password_validation_special"].style.color = BADCOLOR;
    }
    
    if (__testPasswordisAtLeast8Characters(currentpassword.value) &&
        __testPasswordisHasAtLeastOneLowerCaseCharacter(currentpassword.value) &&
        __testPasswordisHasAtLeastOneUpperCaseCharacter(currentpassword.value) &&
        __testPasswordisHasAtLeastOneSpecialCharacter(currentpassword.value)){ 
        currentpassword.style.background = OKCOLOR;
    }
    else{
        currentpassword.style.background = BADCOLOR;
    }

    if (__testPasswordisAtLeast8Characters(newpassword.value) &&
        __testPasswordisHasAtLeastOneLowerCaseCharacter(newpassword.value) &&
        __testPasswordisHasAtLeastOneUpperCaseCharacter(newpassword.value) &&
        __testPasswordisHasAtLeastOneSpecialCharacter(newpassword.value)){ 
        newpassword.style.background = OKCOLOR;
    }
    else{
        newpassword.style.background = BADCOLOR;
    }
    if ( __validateRepeat(repeatnewpassword.value)){
        repeatnewpassword.style.background = OKCOLOR;
    }else{
        repeatnewpassword.style.background = BADCOLOR;
    }
    if (__testPasswordisAtLeast8Characters(repeatnewpassword.value) &&
        __testPasswordisHasAtLeastOneLowerCaseCharacter(repeatnewpassword.value) &&
        __testPasswordisHasAtLeastOneUpperCaseCharacter(repeatnewpassword.value) &&
        __testPasswordisHasAtLeastOneSpecialCharacter(repeatnewpassword.value)){ 
        repeatnewpassword.style.background = OKCOLOR;
    }
    else{
        repeatnewpassword.style.background = BADCOLOR;
    }
    if (__testPasswordisAtLeast8Characters(newpassword.value) &&
        __testPasswordisHasAtLeastOneLowerCaseCharacter(newpassword.value) &&
        __testPasswordisHasAtLeastOneUpperCaseCharacter(newpassword.value) &&
        __testPasswordisHasAtLeastOneSpecialCharacter(newpassword.value) &&
        __validateRepeat(repeatnewpassword.value)){
            change_button.disabled = false
            __changeButtonColors(BUTTONOKCOLOR);
        }
    else{
            change_button.disabled = true
            __changeButtonColors(BADCOLOR);
    }
}


function Submit(){
    let password = document.getElementById('Pass').value;
    let repeatpass = document.getElementById('PassRepeat').value;
    let currentpass = document.getElementById('Current').value;
    if (password == null || repeatpass == null || currentpass == null)
        return;
    fetch('/changepassword', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            password : password,
            repeatpassword : repeatpass,
            currentpassword : currentpass
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


