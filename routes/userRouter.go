package routes

import (
	controller "resturant-management/controllers"
	"resturant-management/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUser())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	
}
