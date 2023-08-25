package route

import (
	todo "github.com/b0nbon1/learn-go/todo/pkg/todos"
	"github.com/gofiber/fiber/v2"
)

 func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Swerv 👋!")
	})
 	v1 := app.Group("/api/v1")
	todo.TodosRoutes(v1)
 }