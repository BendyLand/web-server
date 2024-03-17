package todo

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
)

type Task string
type TodoList []Task

func (todos *TodoList) Loop() {
	var input string
	fmt.Printf("\nEnter a command or type `help`:\n")
	fmt.Scanln(&input)
	switch input {
	case "add":
		todos.CreateTask()
		todos.Loop()
	case "delete":
		todos.RemoveTask()
		todos.Loop()
	case "view":
		todos.DisplayTodoList()
		todos.Loop()
	case "exit":
		fmt.Println("Here is your final Todo List:")
		todos.DisplayTodoList()
		fmt.Println("Goodbye!")
	default:
		fmt.Printf("HELP MENU:\nThe available commands are:\nadd\ndelete\nview\nexit\nhelp\n\n")
		todos.Loop()
	}
}

func Greet() {
	fmt.Print("\nWelcome to the Go Todo List! Let's get started!\n")
}

func (todos *TodoList) RemoveTask() {
	fmt.Println("Which task would you like to delete?")
	todos.DisplayTodoList()
	id := todos.getId()
	var newTodoList []Task
	for i, todo := range *todos {
		if i != id-1 {
			newTodoList = append(newTodoList, todo)
		}
	}
	*todos = newTodoList
}

func (todos *TodoList) DisplayTodoList() {
	fmt.Println("Todo List:")
	for i, todo := range *todos {
		fmt.Printf("%d.) %s", i+1, todo) // the todo variable ends with '\n'
	}
	fmt.Println()
}

func (todos *TodoList) CreateTask() {
	fmt.Printf("\nPlease enter the text for your task:\n")
	reader := bufio.NewReader(os.Stdin)
	body, _ := reader.ReadSlice('\n')
	newTask := Task(body)
	*todos = append(*todos, newTask)
}

func (todos *TodoList) getId() int {
	var input string
	fmt.Scanln(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("\nInvalid input. Please enter a number.\n")
		todos.getId()
	}
	if id > len(*todos) {
		fmt.Printf("\nInvalid input. Please enter a number under: %d\n", len(*todos)+1)
		todos.getId()
	}
	return id
}
