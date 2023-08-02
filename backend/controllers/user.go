package controllers

import (
	"dbo-backend/database"
	"dbo-backend/models"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (ctrl *UserController) Index(c *gin.Context) {
	q := c.Request.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))

	if page == 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(q.Get("limit"))

	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}

	offset := (page - 1) * limit

	var user []models.User

	result := database.Db.Where("deleted_at IS NULL").Order("id ASC").Find(&user)

	count := result.RowsAffected

	result.Limit(limit).Offset(offset).Find(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       user,
		"page":       page,
		"total_page": math.Ceil(float64(count) / float64(limit)),
	})
}

func (ctrl *UserController) Detail(c *gin.Context) {
	userId := c.Param("user_id")

	var user models.User

	result := database.Db.Where("users.id=?", userId).First(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (ctrl *UserController) Create(c *gin.Context) {
	type body struct {
		Email        string `json:"email" binding:"required"`
		Name         string `json:"name" binding:"required"`
		MobileNumber string `json:"mobile_number" binding:"required"`
		Password     string `json:"password" binding:"required"`
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

	result := database.Db.Create(&models.User{
		Email:        b.Email,
		Name:         &b.Name,
		MobileNumber: &b.MobileNumber,
		Password:     b.Password,
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
		"message": "Successfully create user.",
	})
}

func (ctrl *UserController) Edit(c *gin.Context) {
	userId := c.Param("user_id")

	type body struct {
		Email        string `json:"email"`
		Name         string `json:"name"`
		MobileNumber string `json:"mobile_number"`
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

	result := database.Db.Where("id=?", userId).Updates(&models.User{
		Email:        b.Email,
		Name:         &b.Name,
		MobileNumber: &b.MobileNumber,
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
		"message": "Successfully edit user.",
	})
}

func (ctrl *UserController) Delete(c *gin.Context) {
	userId := c.Param("user_id")

	result := database.Db.Model(&models.User{}).Where("id=?", userId).Update("deleted_at", time.Now())
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully delete user.",
	})
}
