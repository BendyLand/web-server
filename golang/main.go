package main

import (
	"golang/actions"
	"golang/handlers"
)

func main() {
	actions.Greet()
	taskManager := handlers.NewTaskManager()
	taskManager.CreateDbTable()
	actions.AddTask(taskManager)
	actions.AddTask(taskManager)
	actions.AddTask(taskManager)
	actions.DeleteTask(taskManager)
	taskManager.Shutdown()
}
