package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-products-api/model"
	"github.com/nadirbasalamah/go-products-api/service"
)

func GetAllProducts(c *fiber.Ctx) error {
	var products []model.Product = service.GetAllProducts()

	return c.JSON(products)
}

func GetProductById(c *fiber.Ctx) error {
	var productId string = c.Params("id")

	var product model.Product = service.GetProductById(productId)

	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	var productInput *model.ProductInput = new(model.ProductInput)

	if err := c.BodyParser(productInput); err != nil {
		c.Status(503).SendString(err.Error())
		return err
	}

	errors := productInput.ValidateStruct()

	if errors != nil {
		return c.Status(400).JSON(errors)
	}

	var createdProduct model.Product = service.CreateProduct(*productInput)

	return c.JSON(createdProduct)
}

func UpdateProduct(c *fiber.Ctx) error {

	var productInput *model.ProductInput = new(model.ProductInput)

	if err := c.BodyParser(productInput); err != nil {
		c.Status(503).SendString(err.Error())
		return err
	}

	errors := productInput.ValidateStruct()

	if errors != nil {
		return c.Status(400).JSON(errors)
	}

	var productId string = c.Params("id")

	var updatedProduct model.Product = service.UpdateProduct(productId, *productInput)

	return c.JSON(updatedProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	var productId string = c.Params("id")

	var result bool = service.DeleteProduct(productId)

	if result {
		c.SendString("Data deleted")
	}

	return nil
}

func parseRequest(c *fiber.Ctx, productInput model.ProductInput) error {
	if err := c.BodyParser(productInput); err != nil {
		return err
	}
	return nil
}

func validateRequest(productInput model.ProductInput) []*model.ErrorResponse {
	errors := productInput.ValidateStruct()

	if errors != nil {
		return errors
	}

	return nil
}
