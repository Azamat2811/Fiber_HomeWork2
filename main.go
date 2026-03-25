package main

import (
	
	"fiber_hw_2/db"
	"fiber_hw_2/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()

	app := fiber.New()

	routes.UserRoutes(app)


	if err := app.Listen(":3000"); err != nil {

	}
}

