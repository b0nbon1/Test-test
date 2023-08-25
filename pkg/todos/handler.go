package todo

import (
	"fmt"

	"github.com/b0nbon1/learn-go/todo/pkg/todos/db"
	initdb "github.com/b0nbon1/learn-go/todo/postgres"
	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	queries := db.New(initdb.DB)

	todos, err := queries.GetAllTodos(c.Context())
	if err != nil {
		fmt.Printf("Error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println(todos)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"name": "Work at home",
		"completed": false,
	})
}

