package routes

import (
	"github.com/amillerrr/htmx-todo-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	todohandler := &handlers.TodoHandler{}
	todogroup := app.Group("/api/v1/todos")
	todogroup.Get("/", todohandler.FetchTodos)
	todogroup.Get("/:id", todohandler.FetchTodo)
	todogroup.Get("/:id", todohandler.MarkDone)
	todogroup.Post("/create", todohandler.CreateTodo)
	todogroup.Post("/:id", todohandler.DeleteTodo)
}
