package services

import (
	"errors"
	"fmt"

	"github.com/amillerrr/htmx-todo-app/db"
	"github.com/amillerrr/htmx-todo-app/models"
	"github.com/google/uuid"
)

// GetAllTodos returns all todo data
func GetAllTodos() ([]models.Todo, error) {
	// create a variable to store todo data
	var todos []models.Todo = []models.Todo{}
	// get all data from the database order by created_at
	rows, err := db.DB.Query("SELECT id, todo, done FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var todo string
		var done bool
		err := rows.Scan(&id, &todo, &done)
		if err != nil {
			return nil, err
		}
		nexttodo := models.Todo{
			ID:   id,
			Todo: todo,
			Done: done,
		}
		todos = append(todos, nexttodo)
	}
	// return all todos from the database
	return todos, err
}

// GetTodoByID returns get todo's data by ID
func GetTodoByID(id string) (models.Todo, error) {
	// create a variable to store todo data
	var todo models.Todo
	todo.ID = id
	if id == "" {
		return todo, errors.New("400. Bad Request")
	}
	// get todo data from the database by ID
	row, err := db.DB.Query("SELECT todo, done FROM todos WHERE id = $1;", id)
	// if the todo data is not found, return an error
	if err != nil {
		return models.Todo{}, errors.New("item not found")
	}
	for row.Next() {
		err := row.Scan(&todo.Todo, &todo.Done)
		if err != nil {
			return models.Todo{}, err
		}
	}
	// return the todo data from the database
	return todo, nil
}

// CreateTodo returns created todo in the storage
func CreateTodo(todo string) error {
	// create a new todo
	var newTodo models.Todo = models.Todo{
		ID:   uuid.New().String(),
		Todo: todo,
		Done: false,
	}
	// validate form value
	switch {
	case newTodo.ID == "":
		return errors.New("400. Bad Request. ID field Invalid")
	case newTodo.Todo == "":
		return errors.New("400. Bad Request. Todo field empty")
	}
	// insert the created todo into the database
	statement := "INSERT INTO todos (id, todo, done) VALUES ($1, $2, $3)"
	_, err := db.DB.Query(statement, newTodo.ID, newTodo.Todo, newTodo.Done)
	if err != nil {
		return errors.New("500. Internal Server Error." + err.Error())
	}
	// return the todo that was created
	return err
}

// MarkDone toggles a todo as completed
func MarkDone(id string) error {
	// get the todo data by ID
	todo, err := GetTodoByID(id)
	// if item is not found, return an error
	if err != nil {
		return errors.New("400. Bad Request")
	}
	_, err = db.DB.Query("UPDATE todos SET done=$2 WHERE id=$1", todo.ID, !todo.Done)
	if err != nil {
		// return error if update is failed
		return errors.New("item completion failed, item not found")
	}
	return err
}

// DeleteTodo returns deletion nothing unless there is an error
func DeleteTodo(id string) error {
	// get the todo data by ID
	todo, err := GetTodoByID(id)
	fmt.Println(todo.ID)
	// if todo is not found, return false
	if err != nil {
		return errors.New("400. Bad Request")
	}
	// delete the todo data
	_, err = db.DB.Exec("DELETE FROM todos WHERE id=$1", todo.ID)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	// return true that the deletion succeeded
	return err
}
