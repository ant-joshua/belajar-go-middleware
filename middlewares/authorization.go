package middlewares

import (
	"belajar-middleware/database"
	"belajar-middleware/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		productID, err := strconv.Atoi(c.Param("productID"))

		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"message": "invalid parameter product id",
				"error":   "Bad Request",
			})
			return
		}

		user := c.MustGet("user").(jwt.MapClaims)
		userID := uint(user["id"].(float64))
		role := user["role"].(string)
		fmt.Printf("%v", user)

		product := models.Product{}

		err = db.Debug().Where("id = ?", productID).First(&product).Error

		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{
				"message": "product not found",
				"error":   "Not Found",
			})
			return
		}

		fmt.Println("product user id", product.UserID, "user id", userID)
		fmt.Println("role", role)

		if product.UserID != userID && role != "admin" {
			c.AbortWithStatusJSON(403, gin.H{
				"message": "you are not authorized to access this product",
				"error":   "Forbidden",
			})
			return
		}

		c.Next()
	}
}
