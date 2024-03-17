package main

import (
	"fmt"
	"golang/todo"
	"net/http"
)

func main() {
    todo.Greet()

	var todos todo.TodoList
    todos.Loop()
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
