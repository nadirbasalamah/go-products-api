package service

import (
	"github.com/nadirbasalamah/go-products-api/model"
	"github.com/nadirbasalamah/go-products-api/repository"
)

func GetAllProducts() []model.Product {
	return repository.GetAllProducts()
}

func GetProductById(id string) model.Product {
	product, _ := repository.GetProductById(id)
	return product
}

func CreateProduct(input model.ProductInput) model.Product {
	return repository.CreateProduct(input)
}

func UpdateProduct(id string, input model.ProductInput) model.Product {
	return repository.UpdateProduct(id, input)
}

func DeleteProduct(id string) bool {
	return repository.DeleteProduct(id)
}
