package cmd

import (
	"tasks/cmd/handlers"

	"github.com/gofiber/fiber/v2"
)

func RoutesInit(app *fiber.App) {
	r := app
	r.Get("/", handlers.Home)
	r.Post("/add-task", handlers.AddTask)

	r.Get("/projects", handlers.Projects)
	r.Post("/add-project", handlers.AddProject)
}
