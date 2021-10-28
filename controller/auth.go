package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-products-api/model"
	"github.com/nadirbasalamah/go-products-api/service"
)

func Register(c *fiber.Ctx) error {
	var userInput *model.UserInput = new(model.UserInput)

	if err := c.BodyParser(userInput); err != nil {
		c.Status(503).SendString(err.Error())
		return err
	}

	errors := userInput.ValidateStruct()

	if errors != nil {
		return c.Status(400).JSON(errors)
	}

	token, err := service.Register(*userInput)

	if err != nil {
		c.Status(500).JSON(err)
		return err
	}

	return c.JSON(fiber.Map{"token": token})
}

func Login(c *fiber.Ctx) error {
	var userInput *model.UserInput = new(model.UserInput)

	if err := c.BodyParser(userInput); err != nil {
		c.Status(503).SendString(err.Error())
		return err
	}

	errors := userInput.ValidateStruct()

	if errors != nil {
		return c.Status(400).JSON(errors)
	}

	token, err := service.Login(*userInput)

	if err != nil {
		c.Status(500).JSON(err)
		return err
	}

	return c.JSON(fiber.Map{"token": token})
}
