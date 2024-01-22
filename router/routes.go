package router

import (
	"github.com/gin-gonic/gin"
	"github.com/matfigueiredo/go-opportunities/handler"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()

	v1 := router.Group("/api/v1")

	{
		v1.GET("/jobs_opening", handler.ShowOpeningHandler)

		v1.POST("/jobs_opening", handler.CreateOpeningHandler)

		v1.DELETE("/jobs_opening", handler.DeleteOpeningHandler)

		v1.PUT("/jobs_opening", handler.UpdateOpeningHandler)

		v1.GET("/jobs_openings", handler.ListOpeningsHandler)
	}
}
