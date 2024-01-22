package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/matfigueiredo/go-opportunities/docs"
	"github.com/matfigueiredo/go-opportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/api/v1")

	{
		v1.GET("/jobs_opening", handler.ShowOpeningHandler)

		v1.POST("/jobs_opening", handler.CreateOpeningHandler)

		v1.DELETE("/jobs_opening", handler.DeleteOpeningHandler)

		v1.PUT("/jobs_opening", handler.UpdateOpeningHandler)

		v1.GET("/jobs_openings", handler.ListOpeningsHandler)
	}

	//Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
