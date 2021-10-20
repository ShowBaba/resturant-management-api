package routes

import (
	"github.com/gin-gonic/gin"

	controller "resturant-management/controllers"
)

func MenuRoutes(incomingRoutes *gin.RouterGroup){
	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id", controller.UpdateMenu())
}