package routes

import (
	"crud/controllers"
	"crud/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", middleware.VerifyToken, controllers.ShowBook)
			books.POST("/", middleware.VerifyToken, controllers.CreateBook)
		}

	}
	return router
}
