package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-products-api/controller"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/products", controller.GetAllProducts)
	app.Get("/api/products/:id", controller.GetProductById)
	app.Post("/api/products", controller.CreateProduct)
	app.Put("/api/products/:id", controller.UpdateProduct)
	app.Delete("/api/products/:id", controller.DeleteProduct)
}
