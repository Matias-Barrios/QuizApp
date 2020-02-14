let username = document.getElementsById('Name')
let email = document.getElementsById('Email')
let password = document.getElementsById('Pass')
let password_check = document.getElementsById('PassRepeat')


function validateEmail(input) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(input).toLowerCase());
}

function userName(input) {
    var re = /^[a-zA-Z][a-zA-Z0-9_-]{5,}$/;
    return re.test(String(input).toLowerCase());
}


function userName(input) {
    var re = /^[a-zA-Z][a-zA-Z0-9_-]{5,}$/;
    return re.test(String(input).toLowerCase());
}

function testPasswordisAtLeast8Characters(input){
    var re = /^.*{8,}$/;
    return re.test(String(input).toLowerCase());
}

function testPasswordisHasAtLeastOneLowerCaseCharacter(input){
    var re = /^(?=.*[a-z])$/;
    return re.test(String(input).toLowerCase());
}

function testPasswordisHasAtLeastOneUpperCaseCharacter(input){
    var re = /^(?=.*[A-Z])$/;
    return re.test(String(input).toLowerCase());
}

function testPasswordisHasAtLeastOneSpecialCharacter(input){
    var re = /^(?=.*[!@#$%^&*])$/;
    return re.test(String(input).toLowerCase());
}


