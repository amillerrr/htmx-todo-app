package main

import (
	"log"
	"net/http"

	"github.com/amillerrr/htmx-todo-app/handlers"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/todos", handlers.Index)
	mux.HandleFunc("/todos/{id}", handlers.MarkDone).Methods("PUT")
	mux.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
	mux.HandleFunc("/todos/create", handlers.CreateTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", mux))
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "/todos", http.StatusSeeOther)
// }
