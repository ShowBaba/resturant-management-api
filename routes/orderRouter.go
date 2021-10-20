package routes

import (
	"github.com/gin-gonic/gin"

	controller "resturant-management/controllers"
)

func OrderRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controller.GetOrder())
	incomingRoutes.POST("/orders", controller.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())
}
