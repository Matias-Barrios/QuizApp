package views

import "text/template"

// ViewIndex :
var ViewIndex = template.Must(template.ParseFiles("./views/view_index/index.html"))
