package middleware

import (
	database "finance-tracker-api/Database"
	models "finance-tracker-api/Models"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(c *gin.Context) {
	// Get
	tokenString, err := c.Cookie("Authorization")
	tokenString = strings.TrimSpace(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized | Token doesn't exist",
		})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized | Invalid token, error parse",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized | Invalid token, token is expired",
			})
			c.Abort()
			return
		}

		var user models.User
		database.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized | Invalid token, user not found",
			})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized | Invalid token, error claims token not found",
		})
		return
	}
}
