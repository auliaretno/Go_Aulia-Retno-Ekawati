package database

import (
	"code_competence/config"
	"code_competence/models"
)

func GetProductsController() (interface{}, error) {
	var products []models.Product

	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductController(productID int) (interface{}, error) {
	var product models.Product
	product.ID = uint(productID)

	if err := config.DB.First(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func CreateProductController(b models.Product) (interface{}, error) {
	err := config.DB.Create(&b).Error

	if err != nil {
		return nil, err
	}

	return b, nil
}

func UpdateProductController(productID uint, b models.Product) (interface{}, error) {
	product := models.Product{}
	product.ID = productID
	if err := config.DB.First(&product).Error; err != nil {
		return nil, err
	}

	product.Merk = b.Merk
	product.Description = b.Description
	product.Stock = b.Stock
	product.Price = b.Price

	if err := config.DB.Save(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func DeleteProductController(productID int) (interface{}, error) {
	err := config.DB.Delete(&models.Product{}, productID).Error

	if err != nil {
		return nil, err
	}
	return productID, nil
}
