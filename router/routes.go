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

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)

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
