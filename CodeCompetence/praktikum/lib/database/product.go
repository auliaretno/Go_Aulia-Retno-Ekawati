package database

import (
	"code_competence/config"
	"code_competence/models"
)

func GetBooksController() (interface{}, error) {
	var books []models.Product

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookController(bookID int) (interface{}, error) {
	var book models.Product
	book.ID = uint(bookID)

	if err := config.DB.First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func CreateBookController(b models.Product) (interface{}, error) {
	err := config.DB.Create(&b).Error

	if err != nil {
		return nil, err
	}

	return b, nil
}

func UpdateBookController(bookID uint, b models.Product) (interface{}, error) {
	book := models.Product{}
	book.ID = bookID
	if err := config.DB.First(&book).Error; err != nil {
		return nil, err
	}

	book.Merk = b.Merk
	book.Description = b.Description
	book.Stock = b.Stock
	book.Price = b.Price

	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBookController(bookID int) (interface{}, error) {
	err := config.DB.Delete(&models.Product{}, bookID).Error

	if err != nil {
		return nil, err
	}
	return bookID, nil
}
