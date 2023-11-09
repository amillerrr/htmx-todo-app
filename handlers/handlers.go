package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/amillerrr/htmx-todo-app/services"
	"github.com/gorilla/mux"
)

func sendTodos(w http.ResponseWriter) {

	todos, err := services.GetAllTodos()
	if err != nil {
		fmt.Println("Could not get all todos from db", err)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	err = tmpl.ExecuteTemplate(w, "Todos", todos)
	if err != nil {
		fmt.Println("Could not execute template", err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {

	todos, err := services.GetAllTodos()
	if err != nil {
		fmt.Println("Could not get all todos from db", err)
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	err = tmpl.Execute(w, todos)
	if err != nil {
		fmt.Println("Could not execute template", err)
	}

}

func MarkDone(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	err := services.MarkDone(id)
	if err != nil {
		fmt.Println("Could not update todo", err)
	}

	sendTodos(w)

}

func CreateTodo(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form", err)
	}

	err = services.CreateTodo(r.FormValue("todo"))
	if err != nil {
		fmt.Println("Could not create todo", err)
	}

	sendTodos(w)

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := services.DeleteTodo(id)
	if err != nil {
		fmt.Println("Could not delete", err)
	}

	sendTodos(w)

}
