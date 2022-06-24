package routes

import (
	"modulo/controllers"

	
	"github.com/labstack/echo/v4"
)

func ConfigRoutes(router *echo.Echo) *echo.Echo {

	// router.GET("/")

	main := router.Group("/api")
	{
		estabelecimentos := main.Group("/estabelecimentos")
		{
			estabelecimentos.GET("/:id", controllers.ShowEstablishment)
			estabelecimentos.GET("/", controllers.ShowEstablishments)
		}
		// lojas := main.Group("lojas")
		// {

		// }
	}

	return router
}