package views

import "text/template"

// ViewInternalServerError :
var ViewInternalServerError = template.Must(template.ParseFiles("static/internalServerError/error.html"))
