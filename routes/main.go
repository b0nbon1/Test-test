package route

import (
	todo "github.com/b0nbon1/learn-go/todo/pkg/todos"
	user "github.com/b0nbon1/learn-go/todo/pkg/users"
	"github.com/gofiber/fiber/v2"
)

//	@BasePath	/api/v1
 func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Swerv 👋!")
	})
 	v1 := app.Group("/api/v1")
	todo.TodosRoutes(v1)
	user.UsersRoutes(v1)
 }