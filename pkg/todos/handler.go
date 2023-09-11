package todo

import (
	"fmt"

	"github.com/b0nbon1/learn-go/todo/pkg/todos/db"
	initdb "github.com/b0nbon1/learn-go/todo/postgres"
	util "github.com/b0nbon1/learn-go/todo/utils"
	"github.com/gofiber/fiber/v2"
)

type ReturnedTodo struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
}

type TodoResponse struct {
	data ReturnedTodo
	message string
}

func mapTodo(todo db.Todo) interface{} {

	return ReturnedTodo{
		ID: todo.ID,
		Name: todo.Name,
		Completed: todo.Completed.Bool,
	}
}

// CreateTodo func creates Todos.
// @Description Be able to create a Todo
// @Summary create a Todo
// @Tags Todo
// @Accept json
// @Produce json
// @Param todo body ReturnedTodo false "Todo"
// @Success 200 {object} TodoResponse
// @Router /api/v1/todos [post]
func CreateTodo(c *fiber.Ctx) error {
	queries := db.New(initdb.DB)

	todos, err := queries.GetAllTodos(c.Context())
	if err != nil {
		fmt.Printf("Error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	newTodos := util.MapValues(todos, mapTodo)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data": newTodos,
	})
}

