package router

import (
	"github.com/gin-gonic/gin"
	"github.com/seven-yu/gin-auto-bind/app/controller"
)

func RouteV1(e *gin.Engine) {
	apiV1 := e.Group("/api/v1")

	{
		apiV1.GET("hello", controller.Get)
	}
}
