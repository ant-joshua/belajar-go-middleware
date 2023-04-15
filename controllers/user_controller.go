package controllers

import (
	"belajar-middleware/database"
	"belajar-middleware/helpers"
	"belajar-middleware/models"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(c)

	user := models.User{}

	// error
	var err error

	if contentType != "application/json" {
		err = c.ShouldBindJSON(&user)
	} else {
		err = c.ShouldBind(&user)
	}

	if err != nil {
		panic(err)
	}

	err = db.Debug().Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to register",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success to register",
		"data": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
	})

}

func UserLogin(c *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(c)

	user := models.User{}

	var err error

	if contentType != "application/json" {
		err = c.ShouldBindJSON(&user)
	} else {
		err = c.ShouldBind(&user)
	}

	if err != nil {
		panic(err)
	}
	password := user.Password

	err = db.Debug().Where("email = ?", user.Email).Take(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Username not found",
			"error":   err.Error(),
		})
		return
	}

	if !helpers.CheckPasswordHash(password, user.Password) {
		c.JSON(400, gin.H{
			"message": "Password not match",
			"error":   "asd",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email, user.Role)

	c.JSON(200, gin.H{
		"message": "Success to register",
		"data": gin.H{
			"access_token": token,
		},
	})

}
