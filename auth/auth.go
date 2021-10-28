package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-products-api/utils"
)

// NOT USED!
func Register() (string, error) {
	token, err := utils.GenerateNewAccessToken()

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetResources(c *fiber.Ctx) (string, error) {
	isValid, err := utils.CheckToken(c)

	if !isValid {
		return "", err
	}

	return "welcome!", nil
}
