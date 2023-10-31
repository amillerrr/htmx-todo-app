package models

import (
	"github.com/jmoiron/sqlx"
)

type TodoStorage struct {
	Conn *sqlx.DB
}

type Todo_DB struct {
	Id   int    `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

func NewTodoStorage(conn *sqlx.DB) *TodoStorage {
	return &TodoStorage{Conn: conn}
}

type NewTodoInput struct {
	Todo string
}

func (s *TodoStorage) CreateTodo(todo NewTodoInput) error {

	statement := "INSERT INTO todos (todo, done) VALUES (?, ?)"

	_, err := s.Conn.Exec(statement, todo.Todo, false)

	return err

}

func (s *TodoStorage) GetAllTodos() ([]Todo_DB, error) {
	todos := []Todo_DB{}
	err := s.Conn.Select(&todos, "SELECT id, todo, done FROM todos")
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (s *TodoStorage) GetTodo(id int) (Todo_DB, error) {
	todo := Todo_DB{}
	todo.Id = id
	err := s.Conn.Get(&todo, "SELECT todo, done FROM todos WHERE id=?", id)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (s *TodoStorage) MarkDone(id int) error {
	_, err := s.Conn.Exec("UPDATE todos SET done=true WHERE id=?", id)
	return err
}

func (s *TodoStorage) Delete(id int) error {
	_, err := s.Conn.Exec("DELETE FROM todos WHERE id=?", id)
	return err
}
