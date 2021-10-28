package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-products-api/controller"
	"github.com/nadirbasalamah/go-products-api/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/products", controller.GetAllProducts)
	app.Get("/api/products/:id", controller.GetProductById)

	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
}

func SetupPrivateRoutes(app *fiber.App) {
	app.Post("/api/products", middleware.JWTProtected(), controller.CreateProduct)
	app.Put("/api/products/:id", middleware.JWTProtected(), controller.UpdateProduct)
	app.Delete("/api/products/:id", middleware.JWTProtected(), controller.DeleteProduct)
}
