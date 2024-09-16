package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasdbr05/gopportunities/handler"
)

func initRoutes(router *gin.Engine) {
	
	handler.InitHandler()

	api := router.Group("/api/")
	{
		api.GET("/opening", handler.ShowOpeningHandler)
		api.POST("/opening", handler.CreateOpeningHandler)
		api.PATCH("/opening", handler.UpdateOpeningHandler)
		api.DELETE("/opening", handler.DeleteOpeningHandler)
		api.GET("/openings", handler.ListOpeningsHandler)
	}
}
