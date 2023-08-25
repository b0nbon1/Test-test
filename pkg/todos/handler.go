package todo

import (
	// "database/sql"
	"database/sql"
	"fmt"

	"github.com/b0nbon1/learn-go/todo/pkg/todos/db"
	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	dbc := c.Locals("db").(*sql.DB)
	queries := db.New(dbc)
	todos, err := queries.GetAllTodos(c.Context())
	if err != nil {
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

