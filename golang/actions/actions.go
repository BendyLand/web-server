package actions

import (
	"bufio"
	"fmt"
	"golang/handlers"
	"log"
	"net/http"
	"os"
	"strconv"
)

func DeleteTask(m *handlers.TaskManager) {
	var input string
	fmt.Println("Which task would you like to delete?")
	m.DisplayTasks()
	fmt.Scan(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Error getting input: ", err)
	}
	m.DeleteTask(id)
}

func AddTask(m *handlers.TaskManager) {
	fmt.Println("Please enter the text for your task:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("There was a problem getting input")
	}
	m.AddTask(input)
}

func Greet() {
	fmt.Printf("\nWelcome to the Go Todo List! Let's get started!\n\n")
}

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
