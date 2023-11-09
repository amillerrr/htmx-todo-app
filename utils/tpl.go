package utils

import "html/template"

var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseFiles("templates/index.html"))
}
