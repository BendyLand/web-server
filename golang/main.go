package main

import (
	"golang/actions"
	"golang/handlers"
)

func main() {
	actions.Greet()
	taskManager := handlers.NewTaskManager()
	taskManager.CreateDbTable()
	actions.Loop(taskManager)
	taskManager.Shutdown()
}
