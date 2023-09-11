package app

import (
	// "database/sql"
	// "fmt"

	_ "github.com/b0nbon1/learn-go/todo/docs"
	initdb "github.com/b0nbon1/learn-go/todo/postgres"
	router "github.com/b0nbon1/learn-go/todo/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

//	@title			Swerv API
//	@version		1.0
//	@description	This is a sample server for Swerv API.
//	@termsOfService	http://swagger.io/terms/
//	@host			localhost:4500
//	@BasePath		/
func SetupAndRunApp() error {
	err := initdb.InitDb("postgres://bon:1234567890@localhost:5432/sqlc_todos")
	if err != nil {
		return err
	}

	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)


	app.Use(recover.New())

	router.SetupRoutes(app)


	app.Listen(":4500")

	return nil
}
