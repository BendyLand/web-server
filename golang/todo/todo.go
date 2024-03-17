package todo

import (
	"fmt"
)

type Task string
type TodoList []Task

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

