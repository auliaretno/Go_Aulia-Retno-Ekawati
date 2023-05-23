package controllers

import (
	"net/http"
	"strconv"
	"code_competence/config"
	"code_competence/models"

	"github.com/labstack/echo/v4"
)

func GetProductsController(c echo.Context) error {
	var products []models.Product

	if err := config.DB.Find(&products).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get data products",
		"data":    products,
	})
}

// get product by id
func GetProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var product models.Product
	if err = config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get product by id",
		"user":    product,
	})
}

// create new product
func CreateProductController(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)

	if err := config.DB.Save(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new product",
		"user":    product,
	})
}

// delete product by id
func DeleteProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var product models.Product
	if err := config.DB.First(&product, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "product not found",
		})
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete product data",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted product data",
		"data":    product,
	})
}

// update product by id
func UpdateProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var book models.Product
	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "User not found",
		})
	}
	c.Bind(&book)
	if err := config.DB.Model(&book).Updates(book).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book data",
		"user":    book,
	})
}
