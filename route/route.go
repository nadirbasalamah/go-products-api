package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-products-api/auth"
	"github.com/nadirbasalamah/go-products-api/controller"
	"github.com/nadirbasalamah/go-products-api/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/products", controller.GetAllProducts)
	app.Get("/api/products/:id", controller.GetProductById)
	app.Post("/api/products", controller.CreateProduct)
	app.Put("/api/products/:id", controller.UpdateProduct)
	app.Delete("/api/products/:id", controller.DeleteProduct)

	//test only, removed soon
	app.Get("/api/register", func(c *fiber.Ctx) error {
		jwtToken, err := auth.Register()
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{"token": jwtToken})
	})
}

func SetupPrivateRoutes(app *fiber.App) {
	// test only, removed soon
	app.Get("/api/resources", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		result, err := auth.GetResources(c)

		if err != nil {
			return c.Status(403).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{"message": result})
	})
}
