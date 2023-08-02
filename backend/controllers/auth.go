package controllers

import (
	"dbo-backend/database"
	"dbo-backend/models"
	"dbo-backend/services"
	"regexp"

	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type AuthController struct{}

func (ctrl *AuthController) Register(c *gin.Context) {
	type body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var b body

	err := c.ShouldBind(&b)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusPreconditionFailed, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	// Email validation
	emailValid := isEmailValid(b.Email)
	if !emailValid {
		c.JSON(http.StatusPreconditionFailed, gin.H{
			"message": "Invalid email format",
		})
		c.Abort()
		return
	}

	// Hashing password
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	fmt.Println(err)
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": err,
	// 	})
	// 	c.Abort()
	// 	return
	// }

	// Create new user
	result := database.Db.Create(&models.User{
		Email:    b.Email,
		Password: b.Password,
	})
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully register.",
	})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	type body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var b body

	if err := c.ShouldBindBodyWith(&b, binding.JSON); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusPreconditionFailed, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	var user models.User

	result := database.Db.Where("email=?", b.Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Wrong email/password",
			})
			c.Abort()
			return
		} else {
			fmt.Println(result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": result.Error.Error(),
			})
			c.Abort()
			return
		}
	}

	if user.DeletedAt != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Your account has been deleted.",
		})
		c.Abort()
		return
	}

	// Comparing the password with the hash
	if user.Password != b.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong email/password",
		})
		c.Abort()
		return
	}

	// Create token
	token, err := services.GenerateToken(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, token)
}

func (ctrl *AuthController) Index(c *gin.Context) {
	type response struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		MobileNumber string `json:"mobile_number"`
		Email        string `json:"email"`
		Password     string `json:"password"`
		IssuedAt     string `json:"issued_at"`
	}

	var r []response

	result := database.Db.Model(models.AccessToken{}).
		Select("distinct on (access_tokens.user_id) access_tokens.id, name, mobile_number, email, password, issued_at").
		Joins("LEFT JOIN users ON access_tokens.user_id=users.id").
		Order("access_tokens.user_id ASC, issued_at DESC").
		Find(&r)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": r,
	})
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
