package service

import (
	"github.com/nadirbasalamah/go-products-api/model"
	"github.com/nadirbasalamah/go-products-api/repository"
)

func Register(userInput model.UserInput) (string, error) {
	return repository.Register(userInput)
}

func Login(userInput model.UserInput) (string, error) {
	return repository.Login(userInput)
}
