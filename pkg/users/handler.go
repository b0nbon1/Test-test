package user

import (
	"fmt"

	"github.com/b0nbon1/learn-go/todo/pkg/users/db"
	initdb "github.com/b0nbon1/learn-go/todo/postgres"
	// util "github.com/b0nbon1/learn-go/todo/utils"
	"github.com/gofiber/fiber/v2"
)

// func mapTodo(todo db.Todo) interface{} {

// 	return struct {
// 		ID int64 `json:"id"`
// 		Name string `json:"name"`
// 		Completed bool `json:"completed"`
// 	}{
// 		ID: todo.ID,
// 		Name: todo.Name,
// 		Completed: todo.Completed.Bool,
// 	}
// }

func CreateUser(c *fiber.Ctx) error {
	queries := db.New(initdb.DB)

	users, err := queries.GetAllUsers(c.Context())
	if err != nil {
		fmt.Printf("Error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println(users);
	// newTodos := util.MapValues(todos, mapTodo)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data": users,
	})
}

