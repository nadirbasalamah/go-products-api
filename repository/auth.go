package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nadirbasalamah/go-products-api/database"
	"github.com/nadirbasalamah/go-products-api/model"
	"github.com/nadirbasalamah/go-products-api/utils"
)

func Register(userInput model.UserInput) (string, error) {
	//TODO: encrypt password
	var user model.User = model.User{
		ID:       uuid.New().String(),
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	database.DB.Create(&user)

	token, err := utils.GenerateNewAccessToken()

	if err != nil {
		return "", err
	}

	return token, nil
}

func Login(userInput model.UserInput) (string, error) {
	var user model.User

	result := database.DB.First(&user, "email = ?", userInput.Email)

	if result.RowsAffected == 0 {
		return "", errors.New("User not found")
	}

	//TODO: encrypt password
	var isMatch bool = user.Password == userInput.Password

	if !isMatch {
		return "", errors.New("Invalid password")
	}

	token, err := utils.GenerateNewAccessToken()

	if err != nil {
		return "", err
	}

	return token, nil
}
