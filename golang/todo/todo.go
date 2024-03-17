package todo

import (
	"fmt"
	"strconv"
)

type Task string
type TodoList []Task

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
		fmt.Printf("%d.) %s\n", i+1, todo)
	}
	fmt.Println()
}

func (todos *TodoList) CreateTask(body string) {
	newTask := Task(body)
	*todos = append(*todos, newTask)
}

func (todos *TodoList) getId() int {
	var input string
	fmt.Scanln(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input")
		todos.getId()
	}
	if id > len(*todos) {
		fmt.Println("Invalid input. Please enter a number under:", len(*todos)+1)
		todos.getId()
	}
	return id
}
