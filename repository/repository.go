package repository

import (
	"github.com/google/uuid"
	"github.com/nadirbasalamah/go-products-api/model"
)

var storage []model.Product = []model.Product{}

func GetAllProducts() []model.Product {
	return storage
}

func GetProductById(id string) (model.Product, int) {
	var product model.Product
	var productIndex int
	for index, item := range storage {
		if item.ID == id {
			product = item
			productIndex = index
			break
		}
	}

	return product, productIndex
}

func CreateProduct(input model.ProductInput) model.Product {
	var product model.Product = model.Product{
		ID:          uuid.New().String(),
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	}

	storage = append(storage, product)

	return product
}

func UpdateProduct(id string, input model.ProductInput) model.Product {
	product, index := GetProductById(id)

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock

	storage[index] = product

	return product
}

func DeleteProduct(id string) bool {
	var afterDeleted []model.Product = []model.Product{}
	for _, item := range storage {
		if item.ID != id {
			afterDeleted = append(afterDeleted, item)
		}
	}

	storage = afterDeleted

	return true
}
