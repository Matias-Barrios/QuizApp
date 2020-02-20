package views

import "text/template"

var funcs = template.FuncMap{
	"intRange": func(start, end int) []int {
		n := end - start + 1
		result := make([]int, n)
		for i := 0; i < n; i++ {
			result[i] = start + i
		}
		return result
	},
	"mod": func(i, j int) int { return i % j },
	"increment": func(i, j int) int {
		return i + j
	},
	"multiply": func(i, m int) int {
		return i * m
	},
}

// ViewIndex :
var ViewIndex = template.Must(template.New("index.html").Funcs(funcs).ParseFiles("static/indexV/index.html"))
