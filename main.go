package main

import (
	"github.com/Prameesh-P/todo_list/controllers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/todolist", controllers.TodoList)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
