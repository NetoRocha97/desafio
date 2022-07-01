package routes

import (
	"modulo/controllers"

	"github.com/labstack/echo/v4"
)

func ConfigRoutes(router *echo.Echo) *echo.Echo {

	// router.GET("/")

	main := router.Group("/api")
	{
		establishments := main.Group("/estabelecimentos")
		{
			establishments.GET("/:id", controllers.ShowEstablishment)
			establishments.GET("/", controllers.ShowEstablishments)
			establishments.POST("/", controllers.CreateEstablishment)
			establishments.PUT("/", controllers.UpdateEstablishment)
			establishments.DELETE("/:id", controllers.DeleteEstablishment)
			establishments.GET("/:id/lojas", controllers.ShowStoresByEstablishment)
		}
		stores := main.Group("/lojas")
		{
			stores.GET("/:id", controllers.ShowStore)
			stores.GET("/", controllers.ShowStores)
			stores.POST("/", controllers.CreateStore)
			stores.PUT("/", controllers.UpdateStore)
			stores.DELETE("/:id", controllers.DeleteStore)
		}
	}

	return router
}
