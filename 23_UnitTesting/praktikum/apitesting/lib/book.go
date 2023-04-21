package database

import (
	"tugass/config"
	"tugass/models"
)

func GetBooksController() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookController(bookID int) (interface{}, error) {
	var book models.Book
	book.ID = uint(bookID)

	if err := config.DB.First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func CreateBookController(b models.Book) (interface{}, error) {
	err := config.DB.Create(&b).Error

	if err != nil {
		return nil, err
	}

	return b, nil
}

func UpdateBookController(bookID uint, b models.Book) (interface{}, error) {
	book := models.Book{}
	book.ID = bookID
	if err := config.DB.First(&book).Error; err != nil {
		return nil, err
	}

	book.Judul = b.Judul
	book.Penulis = b.Penulis
	book.Penerbit = b.Penerbit

	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBookController(bookID int) (interface{}, error) {
	err := config.DB.Delete(&models.Book{}, bookID).Error

	if err != nil {
		return nil, err
	}
	return bookID, nil
}
