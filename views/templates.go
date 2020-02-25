package views

import "text/template"

// ViewChangePassword :
var ViewChangePassword = template.Must(template.ParseFiles("static/changepassword/change.html"))
