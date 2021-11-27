package handlers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/ilyasssan2/golangGin/controllers"
)

func userHandlers(group *gin.RouterGroup) {
	userGroup := group.Group("/user")
	{
		userGroup.POST("/login", controllers.Login)
		userGroup.POST("/register", controllers.Register)
	}
}
