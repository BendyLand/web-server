package main

import (
    "os"
    "log"
	"fmt"
    "bufio"
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

func addTask(m *handlers.TaskManager) {
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal("There was a problem getting input")
    }
    m.AddTask(input)
}

func greet() {
	fmt.Printf("\nWelcome to the Go Todo List! Let's get started!\n\n")
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}