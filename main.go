package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/go-auth.git/database"
	"github.com/syahlan1/go-auth.git/routes"
)

func main() {
	database.Connect()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
