package routes

import (
	"fiber_hw_2/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {

	users := app.Group("/users")

	users.Post("/", handlers.CreateUsers)
	users.Get("/", handlers.GetUsers)
	users.Get("/:id", handlers.GetUsersByID)
	users.Patch("/:id", handlers.UpdateUser)
	users.Put("/:id", handlers.FullUpdate)
	users.Delete("/:id", handlers.DeleteUser)
	app.Get("/stats", handlers.GetUserStats)
	app.Get("/users", handlers.GetUsersByName)
	app.Get("/users", handlers.GetUsersPagination)

}
