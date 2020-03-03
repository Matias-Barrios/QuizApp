package views

import (
	"math"
	"text/template"
)

var funcs = template.FuncMap{
	"intRange": func(start, end int) []int {
		n := end - start + 1
		result := make([]int, n)
		for i := 0; i < n; i++ {
			result[i] = start + i
		}
		return result
	},
	"paging": func(total, pagesize int) int {
		return int(math.Ceil(float64(total) / float64(pagesize)))
	},
	"lastpage": func(total, pagesize int) int {
		return int(math.Floor(float64(total-1)/float64(pagesize))) * 10
	},
	"increment": func(i, j int) int {
		return i + j
	},
	"multiply": func(i, m int) int {
		return i * m
	},
}

// ViewIndex :
var ViewIndex = template.Must(template.New("index.html").Funcs(funcs).ParseFiles("static/indexV/index.html"))
