package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-products-api/utils"
)

func Register() (string, error) {
	token, err := utils.GenerateNewAccessToken()

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetResources(c *fiber.Ctx) (string, error) {
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return "", err
	}

	expires := claims.Expires

	if now > expires {
		return "", err
	}

	return "welcome!", nil
}
