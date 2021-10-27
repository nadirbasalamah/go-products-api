package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-products-api/database"
	"github.com/nadirbasalamah/go-products-api/route"
)

func main() {
	app := fiber.New()

	database.InitDatabase()

	route.SetupRoutes(app)

	app.Listen(":3000")
}
