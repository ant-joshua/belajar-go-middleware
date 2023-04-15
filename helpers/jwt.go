package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

var secretKey = "secret"

func GenerateToken(id uint, email string, role string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
	})

	tokenString, _ := token.SignedString([]byte(secretKey))

	return tokenString
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("unauthorized")
	headerToken := c.GetHeader("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer ")
	if !bearer {
		return nil, errResponse
	}

	tokenString := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errResponse
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errResponse
	}

}
