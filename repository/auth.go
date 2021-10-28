package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nadirbasalamah/go-products-api/database"
	"github.com/nadirbasalamah/go-products-api/model"
	"github.com/nadirbasalamah/go-products-api/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(userInput model.UserInput) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	var user model.User = model.User{
		ID:       uuid.New().String(),
		Email:    userInput.Email,
		Password: string(password),
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

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))

	if err != nil {
		return "", errors.New("Invalid password")
	}

	token, err := utils.GenerateNewAccessToken()

	if err != nil {
		return "", err
	}

	return token, nil
}
