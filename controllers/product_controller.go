package controllers

import (
	"belajar-middleware/database"
	"belajar-middleware/helpers"
	"belajar-middleware/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(c)

	product := models.Product{}
	user := c.MustGet("user").(jwt.MapClaims)

	userID := uint(user["id"].(float64))

	// error
	var err error

	if contentType != "application/json" {
		err = c.ShouldBindJSON(&product)
	} else {
		err = c.ShouldBind(&product)
	}

	if err != nil {
		panic(err)
	}

	product.UserID = userID

	err = db.Debug().Create(&product).Error

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to register",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success to register",
		"data":    product,
	})

}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(c)

	product := models.Product{}
	user := c.MustGet("user").(jwt.MapClaims)

	userID := uint(user["id"].(float64))

	// error
	var err error

	productID, err := strconv.Atoi(c.Param("productID"))

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Parameter",
			"error":   err.Error(),
		})
	}

	if contentType != "application/json" {
		err = c.ShouldBindJSON(&product)
	} else {
		err = c.ShouldBind(&product)
	}

	if err != nil {
		panic(err)
	}

	product.ID = uint(productID)
	product.UserID = userID

	err = db.Debug().Where("id = ?", productID).Updates(models.Product{
		Title:       product.Title,
		Description: product.Description,
	}).Error

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to register",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success to register",
		"data":    product,
	})

}
