package controllers

import (
	"code_competence/config"
	database "code_competence/lib/database"
	"code_competence/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetItemsController(c echo.Context) error {
	var items []models.Item

	if err := config.DB.Find(&items).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get data items",
		"data":    items,
	})
}

// get item by id
func GetItemController(c echo.Context) error {
	User, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	item, err := database.GetItemController(uint(User))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":  "success get item",
		"items": item,
	})
}

// create new item
func CreateItemController(c echo.Context) error {
	item := models.Item{}
	c.Bind(&item)

	if err := config.DB.Save(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new item",
		"user":    item,
	})
}

// delete item by id
func DeleteItemController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var item models.Item
	if err := config.DB.First(&item, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "item not found",
		})
	}

	if err := config.DB.Delete(&item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete item data",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted item data",
		"data":    item,
	})
}

// update item by id
func UpdateItemController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var item models.Item
	if err := config.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "Item not found",
		})
	}
	c.Bind(&item)
	if err := config.DB.Model(&item).Updates(item).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update item data",
		"user":    item,
	})
}
