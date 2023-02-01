package main

import (
	"github.com/Prameesh-P/todo_list/controllers"
	"net/http"
)

type Todo struct {
	Item string
	Done bool
}
type PageData struct {
	Title string
	Todos []Todo
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/todolist", controllers.Todo)
}
