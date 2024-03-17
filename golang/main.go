package main

import (
	"fmt"
	"golang/todo"
	"net/http"
)

func main() {
	var todos todo.TodoList
	todos.CreateTask("This is my first task")
	todos.CreateTask("This is my second task")
	todos.CreateTask("This is my third task")
	todos.CreateTask("This is my fourth task")
	todos.CreateTask("This is my fifth task")
	todos.DisplayTodoList()
	todos.RemoveTask()
	todos.DisplayTodoList()
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
