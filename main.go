package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"resturant-management/database"
	// "resturant-management/middleware"
	"resturant-management/routes"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}

	router := gin.New()
	router.Use(gin.Logger())
	// router.Use(middleware.Authentication())
	
	v1 := router.Group("/api/v1")
	routes.AuthRoutes(v1)
	routes.UserRoutes(v1)
	routes.FoodRoutes(v1)
	routes.MenuRoutes(v1)
	routes.TableRoutes(v1)
	routes.OrderRoutes(v1)
	routes.OrderItemRoutes(v1)
	routes.InvoiceRoutes(v1)

	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "Home route, see documentation for proper routing!"})
	})

	router.Run(":" + port)
}
