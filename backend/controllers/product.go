package controllers

import (
	"dbo-backend/database"
	"dbo-backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

func (ctrl *ProductController) Index(c *gin.Context) {
	var product []models.Product

	result := database.Db.Find(&product)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}
