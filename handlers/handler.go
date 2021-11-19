package handlers

import "github.com/gin-gonic/gin"

func RoutHandler(group *gin.RouterGroup) {
	userHandlers(group)
	helloHandlers(group)
}
