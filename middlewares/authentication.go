package middlewares

import (
	"belajar-middleware/helpers"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error":   "unauthorized",
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		c.Set("user", verifyToken)
		c.Next()
	}
}
