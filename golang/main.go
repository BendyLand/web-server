package main

import (
	"fmt"
	"golang/handlers"
	"net/http"
)

func main() {
	greet()
	taskManager := handlers.NewTaskManager()
	taskManager.CreateDbTable()
	taskManager.AddTask("This is a test")
	taskManager.AddTask("This is another test")
	taskManager.AddTask("This is a third test")
	taskManager.AddTask("This is a fourth test")
    taskManager.Shutdown()
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func greet() {
	fmt.Printf("\nWelcome to the Go Todo List! Let's get started!\n\n")
}
