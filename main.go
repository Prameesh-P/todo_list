package main

import (
	"html/template"
	"log"
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

func TodoList(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "TODO list",
		Todos: []Todo{
			{Item: "Learn Go  ", Done: true},
			{Item: "Learn flutter", Done: true},
			{Item: "Learn React", Done: true},
		},
	}
	tpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todolist", TodoList)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
