package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-products-api/route"
)

func main() {
	app := fiber.New()

	route.SetupRoutes(app)

	app.Listen(":3000")
}
