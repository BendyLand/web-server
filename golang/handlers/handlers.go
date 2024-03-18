package handlers

import (
	"database/sql"
	"fmt"
	"golang/todo"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type TaskManager struct {
	mutex sync.Mutex
	todos todo.TodoList
	db    *sql.DB
}

func (m *TaskManager) DisplayTasks() {
	fmt.Println("Todo List:")
	res, err := m.db.Query("SELECT * FROM tasks;")
	defer res.Close()
	if err != nil {
		log.Fatal("There was a problem retrieving your tasks")
	}
	for res.Next() {
		var id int
		var body string
		err := res.Scan(&id, &body)
		if err != nil {
			fmt.Println("There was a problem getting this task")
			continue
		}
		fmt.Printf("%d.) %s", id, body) // each body variable already ends in '\n'
	}
}

func (m *TaskManager) DeleteTask(id int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	stmt, err := m.db.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		log.Fatal("There was a problem preparing your query")
	}
	_, err1 := stmt.Exec(id)
	if err1 != nil {
		log.Fatal("There was a problem executing your query")
	}
	fmt.Println("Task deleted successfully!")
}

func (m *TaskManager) Shutdown() {
	if m.db != nil {
		err := m.db.Close()
		if err != nil {
			log.Fatal("There was a problem closing the database connection: ", err)
		}
		fmt.Println("Database connection closed successfully!")
	}
}

func (m *TaskManager) AddTask(body string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	task := todo.Task(body)
	m.todos = append(m.todos, task)
	m.updateDatabase()
}

func (m *TaskManager) CreateDbTable() {
	_, err2 := m.db.Exec(
		`CREATE TABLE 
			IF NOT EXISTS tasks(
				id INTEGER PRIMARY KEY AUTOINCREMENT, 
				body TEXT
			);
		`,
	)
	if err2 != nil {
		log.Fatal("There was a problem executing your query: ", err2)
	}
	fmt.Println("Table created successfully!")
}

func NewTaskManager() *TaskManager {
	newDb := initializeDb()
	return &TaskManager{
		todos: make(todo.TodoList, 0),
		db:    newDb,
	}
}

func (m *TaskManager) nextTask() todo.Task {
	i := len(m.todos) - 1
	return m.todos[i]
}

func initializeDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatal("There was a problem opening the database: ", err)
	}
	return db
}

func (m *TaskManager) updateDatabase() {
	stmt, err := m.db.Prepare("INSERT INTO tasks(body) VALUES (?)")
	if err != nil {
		log.Fatal("There was a problem preparing your query: ", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(m.nextTask())
	if err != nil {
		log.Fatal("There was a problem executing your query: ", err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		log.Fatal("There was a problem getting the rows affected: ", err)
	}
	fmt.Printf("%d rows affected\n", affected)
}
