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
	
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "Home route, see documentation for proper routing!"})
	})

	router.Run(":" + port)
}
