package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"dbo-backend/database"
	"dbo-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func Auth(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := strings.Replace(c.Request.Header["Authorization"][0], "Bearer ", "", -1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}

		return []byte("test123"), nil
	})
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		fmt.Println("claims failed")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var accessToken models.AccessToken

	result := database.Db.Where("id=? AND is_active=TRUE AND expired_at >= NOW()", claims["jti"]).First(&accessToken)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Access token failed",
			})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": result.Error.Error(),
			})
			return
		}
	}

	var data models.User

	result = database.Db.Where("id = ?", claims["sub"]).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println(result.Error.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Access token failed",
			})
			c.Abort()
			return
		} else {
			fmt.Println(result.Error.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": result.Error.Error(),
			})
			c.Abort()
			return
		}
	}

	c.Set("user", data)
	c.Next()
}
