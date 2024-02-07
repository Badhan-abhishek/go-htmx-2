package main

import (
	"tasks/cmd"
	"tasks/cmd/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	models.ConnectDb()

	engine := html.New("./views", ".html")

	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/assets", "./public")

	cmd.RoutesInit(app)

	app.Listen(":3000")
}
