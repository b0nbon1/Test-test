package app

import (
	"database/sql"

	router "github.com/b0nbon1/learn-go/todo/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func SetupAndRunApp() error {
	connectionString := "user=bon password=1234567890 dbname=sqlc_todos port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	router.SetupRoutes(app)

	app.Listen(":4500")

	return nil
}
