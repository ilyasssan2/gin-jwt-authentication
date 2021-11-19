package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasssan2/golangGin/controllers"
	"github.com/ilyasssan2/golangGin/middlewares"
)

func helloHandlers(group *gin.RouterGroup) {
	group.GET("/", middlewares.IsAuth, controllers.Hello)
}
