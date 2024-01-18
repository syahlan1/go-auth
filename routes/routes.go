package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/syahlan1/go-auth.git/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
