package app

import (
	// "database/sql"
	// "fmt"

	initdb "github.com/b0nbon1/learn-go/todo/postgres"
	router "github.com/b0nbon1/learn-go/todo/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupAndRunApp() error {
	err := initdb.InitDb("postgres://bon:1234567890@localhost:5432/sqlc_todos")
	if err != nil {
		return err
	}

	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":4500")

	return nil
}
