package utils

import (
	"html/template"
	"io"
)

//func Render(w io.Writer, name string, files []string, context interface{}) error {
//	tpl := template.Must(template.New(name).ParseFiles(files...))
//	return tpl.Execute(w, context)
//}

func Render(w io.Writer, name string, files []string, context interface{}) error {
	tpl := template.Must(template.ParseFiles(files...))
	return tpl.ExecuteTemplate(w, name, context)
}
