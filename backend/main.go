package main

import (
	"dbo-backend/database"
	"dbo-backend/routers"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Activating: %s\n", os.Getenv("APP_NAME"))

	database.MakeConnection()

	r := gin.Default()

	r.GET("/check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": os.Getenv("APP_NAME") + " server running..."})
	})

	routers.Api(r)

	r.Run(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))

	fmt.Printf("Listening and serving HTTP on %s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
}
