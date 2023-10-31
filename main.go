package main

import (
	"fmt"

	"github.com/amillerrr/htmx-todo-app/config"
	"github.com/amillerrr/htmx-todo-app/db"
	"github.com/amillerrr/htmx-todo-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

const DEFAULT_PORT = "3000"

func NewFiberApp() *fiber.App {
	// Create a new engine
	engine := html.New("./templates", ".html")

	var app *fiber.App = fiber.New(fiber.Config{
		Views: engine,
	})

	routes.SetupRoutes(app)
	return app
}

func main() {
	// create a fiber application
	var app *fiber.App = NewFiberApp()

	db.CreateMySqlConnection(config.GetValue("DB_NAME"))

	// var PORT string = os.Getenv("PORT")

	// if PORT == "" {
	// 	PORT = DEFAULT_PORT
	// }

	app.Listen(fmt.Sprintf(":%s", DEFAULT_PORT))

}
