package routers

import (
	"assignment2/controllers"

	"github.com/gin-gonic/gin"
)

func InitApiRoutes() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api")
	{
		orders := apiGroup.Group("/orders")
		{
			orders.GET("/", controllers.ShowOrder)
			orders.POST("/", controllers.NewOrder)
			orders.DELETE("/:orderId", controllers.DeleteOrder)
			orders.PUT("/:orderId", controllers.UpdateOrder)
			orders.GET("/:orderId", controllers.GetOrder)
		}
	}

	return router
}
