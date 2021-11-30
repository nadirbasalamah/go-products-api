package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-products-api/route"
)

func main() {
	app := fiber.New()

	route.SetupRoutes(app)

	fmt.Println("server started")

	app.Listen(":3000")
}
