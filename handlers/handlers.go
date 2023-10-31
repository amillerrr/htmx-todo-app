package handlers

import (
	"strconv"

	"github.com/amillerrr/htmx-todo-app/models"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	Storage *models.TodoStorage
}

func NewTodoHandler(storage *models.TodoStorage) *TodoHandler {
	return &TodoHandler{Storage: storage}
}

type createTodoRequest struct {
	Todo string `json:"todo"`
}

type createTodoResponse struct {
	// Id int `json:"id"`
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	// get the request body
	var body createTodoRequest
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	// create the todo
	err = h.Storage.CreateTodo(models.NewTodoInput{
		Todo: body.Todo,
	})
	if err != nil {
		return err
	}

	// send the id
	resp := createTodoResponse{}

	return c.JSON(resp)
}

type fetchOneTodoReponse struct {
	Todo models.Todo_DB `json:"todo"`
}

func (h *TodoHandler) FetchTodo(c *fiber.Ctx) error {
	// get the todo
	todos := models.Todo_DB{}
	todo, err := h.Storage.GetTodo(todos.Id)
	if err != nil {
		return err
	}

	// send response
	resp := fetchOneTodoReponse{
		Todo: todo,
	}

	return c.JSON(resp)
}

type fetchTodosRepsonse struct {
	Todos []models.Todo_DB `json:"todos"`
}

func (h *TodoHandler) FetchTodos(c *fiber.Ctx) error {
	// get the todos
	todos, err := h.Storage.GetAllTodos()
	if err != nil {
		return err
	}

	// send Response
	resp := fetchTodosRepsonse{
		Todos: todos,
	}
	return c.JSON(resp)
}

type basicResponse struct {
	Success bool `json:"success"`
}

func (h *TodoHandler) MarkDone(c *fiber.Ctx) error {
	//complete todo
	todos := models.Todo_DB{}
	err := h.Storage.MarkDone(todos.Id)
	if err != nil {
		return err
	}

	// send Response
	resp := basicResponse{
		Success: true,
	}

	return c.JSON(resp)
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	// get the id
	id := c.Params("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	// complete todo
	err = h.Storage.Delete(aid)
	if err != nil {
		return err
	}
	// send response
	resp := basicResponse{
		Success: true,
	}
	return c.JSON(resp)
}
