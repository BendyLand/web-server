package todo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task string
type TodoList []Task

func (todos *TodoList) Loop() {
Loop:
	for {
		var input string
		fmt.Printf("\nEnter a command or type `help`:\n")
		fmt.Scan(&input)
		switch input {
		case "add":
			todos.CreateTask()
		case "delete":
			todos.RemoveTask()
		case "view":
			todos.DisplayTodoList()
		case "exit":
			fmt.Println("Here is your final Todo List:")
			todos.DisplayTodoList()
			fmt.Println("Goodbye!")
			break Loop
		default:
			fmt.Printf("HELP MENU:\nThe available commands are:\nadd\ndelete\nview\nexit\nhelp\n\n")
		}
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
		if i == id {
			fmt.Println("Task deleted successfully!")
			continue
		}
		newTodoList = append(newTodoList, todo)
	}
	*todos = newTodoList
}

func (todos *TodoList) DisplayTodoList() {
	fmt.Println("Todo List:")
	for i, todo := range *todos {
		fmt.Printf("%d.) %s", i+1, todo) // the todo variable already ends with '\n'
	}
	fmt.Println()
}

func (todos *TodoList) CreateTask() {
	fmt.Printf("\nPlease enter the text for your task:\n")
	reader := bufio.NewReader(os.Stdin)
	body, err := reader.ReadSlice('\n')
	if err != nil {
		fmt.Println("There was a problem saving your task, please try again.")
		todos.CreateTask()
	}
	newTask := Task(body)
	*todos = append(*todos, newTask)
	fmt.Println("Task saved successfully!")
}

func (todos *TodoList) getId() int {
	length := len(*todos)
Loop:
	for {
		var input string
		fmt.Scanln(&input)
		id, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("\nInvalid input. Please enter a number.\n")
			continue Loop
		}
		if id > length {
			fmt.Printf("\nInvalid input. Please enter a number under: %d\n", len(*todos)+1)
			continue Loop
		}
		return id - 1
	}
}
