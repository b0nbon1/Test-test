package todo

import "github.com/gofiber/fiber/v2"

func TodosRoutes(router fiber.Router) {
	// router.Get("/todos", getAllTodos)
	// router.Get("/todos/:id", getTodoById)
	router.Post("/todos", CreateTodo)
	// router.Delete("/todos/:id", deleteTodoById)
	// router.Patch("/todos/:id", updateTodoById)
}