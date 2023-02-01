package controllers

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

type Todo struct {
	Item string
	Done bool
}
type PageData struct {
	Title string
	Todos []Todo
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/temp.gohtml"))
}

func TodoList(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "TODO list",
		Todos: []Todo{
			{Item: "GO Is good ", Done: true},
			{Item: "Prameesh is bad", Done: false},
			{Item: "SUN is star..", Done: true},
		},
	}
	tpl.Execute(w, data)
}
