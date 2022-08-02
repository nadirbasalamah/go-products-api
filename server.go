package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-products-api/config"
	"github.com/nadirbasalamah/go-products-api/database"
	"github.com/nadirbasalamah/go-products-api/route"
)

func newFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	route.SetupRoutes(app)

	route.SetupPrivateRoutes(app)

	return app
}

func main() {
	app := newFiberApp()

	database.InitDatabase(config.Config("DB_NAME"))

	app.Listen(":3000")
}
