package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/siralfbaez/go-mvc/controllers"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.PostsIndex)
}
