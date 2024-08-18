package main

import (
	"fmt"
	routes "go-auth-jwt/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{"success": "Access granted by API 1"},
		)
	})

	router.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{"success": "Access granted by API 2"},
		)
	})

	router.Run(":" + port)
	fmt.Println("Server running on port: ", port)
}
