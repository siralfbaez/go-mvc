package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/siralfbaez/go-mvc/controllers"
	"github.com/siralfbaez/go-mvc/initializers"
	"github.com/siralfbaez/go-mvc/middleware"
	"github.com/siralfbaez/go-mvc/routes"
	"os"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnToDb()
	initializers.SyncDB()
}

func main() {
	// Load templates
	engine := html.New("./views", ".html")

	// Setup app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//Routes
	app.Get("/", controllers.PostsIndex)
	app.Use(middleware.RequireAuth)
	routes.Routes(app)

	// Configure app
	app.Static("/", "./public")

	// Start the app
	err := app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		return
	}

}
