package routers

import (
	"dbo-backend/controllers"
	"dbo-backend/middlewares"

	"github.com/gin-gonic/gin"
)

var (
	authController    controllers.AuthController    = controllers.AuthController{}
	userController    controllers.UserController    = controllers.UserController{}
	productController controllers.ProductController = controllers.ProductController{}
	orderController   controllers.OrderController   = controllers.OrderController{}
)

func Api(r *gin.Engine) {
	api := r.Group("api")
	{
		auth := api.Group("auth")
		{
			auth.GET("/", authController.Index)
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}
		user := api.Group("users")
		{
			user.GET("", userController.Index)
			user.GET("/:user_id", userController.Detail)
			user.POST("/create", userController.Create)
			user.PUT("/edit/:user_id", userController.Edit)
			user.DELETE("/delete/:user_id", userController.Delete)
		}
		order := api.Group("orders", middlewares.Auth)
		{
			order.GET("", orderController.Index)
			order.GET("/:order_id", orderController.Detail)
			order.POST("/create", orderController.Create)
			order.PUT("/edit/:order_id", orderController.Edit)
			order.DELETE("/delete/:order_id", orderController.Delete)
		}
		product := api.Group("products", middlewares.Auth)
		{
			product.GET("", productController.Index)
		}
	}
}
