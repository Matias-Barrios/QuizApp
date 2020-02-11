package views

import "text/template"

// View404 :
var View404 = template.Must(template.ParseFiles("static/notFound404/404.html"))
