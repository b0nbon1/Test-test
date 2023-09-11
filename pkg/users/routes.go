package user

import (
	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(router fiber.Router) {
	// router.Get("/todos", getAllTodos)
	// router.Get("/todos/:id", getTodoById)
	router.Post("/users", CreateUser)
	// router.Delete("/todos/:id", deleteTodoById)
	// router.Patch("/todos/:id", updateTodoById)
}
