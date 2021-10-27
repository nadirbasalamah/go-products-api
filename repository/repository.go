package repository

import (
	"github.com/google/uuid"
	"github.com/nadirbasalamah/go-products-api/database"
	"github.com/nadirbasalamah/go-products-api/model"
)

var storage []model.Product = []model.Product{}

func GetAllProducts() []model.Product {
	var products []model.Product

	database.DB.Find(&products)

	return products
}

func GetProductById(id string) (model.Product, int64) {
	var product model.Product

	result := database.DB.First(&product, "id = ?", id)

	if result.RowsAffected == 0 {
		return model.Product{}, 0
	}

	return product, result.RowsAffected
}

func CreateProduct(input model.ProductInput) model.Product {
	var product model.Product = model.Product{
		ID:          uuid.New().String(),
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	}

	database.DB.Create(&product)

	return product
}

func UpdateProduct(id string, input model.ProductInput) model.Product {
	product, rows := GetProductById(id)

	if rows != 0 {
		product.Name = input.Name
		product.Description = input.Description
		product.Price = input.Price
		product.Stock = input.Stock

		database.DB.Save(&product)

		return product
	}

	return model.Product{}
}

func DeleteProduct(id string) bool {
	product, rows := GetProductById(id)

	if rows != 0 {
		database.DB.Delete(&product)
		return true
	}

	return false
}
