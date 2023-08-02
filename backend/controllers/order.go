package controllers

import (
	"dbo-backend/database"
	"dbo-backend/models"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct{}

func (ctrl *OrderController) Index(c *gin.Context) {
	user := c.MustGet("user").(models.User)

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

	type response struct {
		ID        int       `json:"id"`
		Quantity  int       `json:"quantity"`
		Name      string    `json:"name"`
		Price     int       `json:"price"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	var r []response

	result := database.Db.
		Model(models.Order{}).
		Select("orders.id, quantity, products.name, products.price,orders.created_at, orders.updated_at").
		Joins("LEFT JOIN products ON orders.product_id=products.id").
		Where("user_id=? AND orders.deleted_at IS NULL", user.ID).
		Order("orders.created_at DESC").
		Find(&r)

	count := result.RowsAffected

	result.Limit(limit).Offset(offset).Find(&r)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       r,
		"page":       page,
		"total_page": math.Ceil(float64(count) / float64(limit)),
	})
}

func (ctrl *OrderController) Detail(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	orderId := c.Param("order_id")

	type response struct {
		ID           int       `json:"id"`
		Quantity     int       `json:"quantity"`
		Name         string    `json:"name"`
		Price        int       `json:"price"`
		Total        int       `json:"total"`
		SupplierName string    `json:"supplier_name"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	var r []response

	err := c.ShouldBind(&r)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusPreconditionFailed, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	result := database.Db.
		Model(models.Order{}).
		Select("orders.id, quantity, products.name, products.price, products.price*quantity AS total, suppliers.name AS supplier_name, orders.created_at, orders.updated_at").
		Joins("LEFT JOIN products ON orders.product_id=products.id").
		Joins("LEFT JOIN suppliers ON products.supplier_id=suppliers.id").
		Where("user_id=? AND orders.id=? AND orders.deleted_at IS NULL", user.ID, orderId).
		First(&r)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		c.Abort()
		return
	}

	if len(r) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found.",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": r,
	})
}

func (ctrl *OrderController) Create(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	type body struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
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

	var product models.Product

	result := database.Db.Where("id=?", b.ProductID).First(&product)
	if result.Error != nil {
		fmt.Println(result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Product not found.",
			})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": result.Error.Error(),
			})
			c.Abort()
			return
		}
	}

	result = database.Db.Create(&models.Order{
		UserID:    int(user.ID),
		ProductID: b.ProductID,
		Quantity:  b.Quantity,
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
		"message": "Succesfully create order.",
	})
}

func (ctrl *OrderController) Edit(c *gin.Context) {
	orderId := c.Param("order_id")

	type body struct {
		Quantity int `json:"quantity"`
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

	result := database.Db.Model(&models.Order{}).Where("id=?", orderId).Update("quantity", b.Quantity)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully edit order.",
	})
}

func (ctrl *OrderController) Delete(c *gin.Context) {
	orderId := c.Param("order_id")

	result := database.Db.Model(&models.Order{}).Where("id=?", orderId).Update("deleted_at", time.Now())
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully delete order.",
	})
}
