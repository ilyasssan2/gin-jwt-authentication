package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ilyasssan2/golangGin/handlers"
	"github.com/ilyasssan2/golangGin/utils"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cancel := utils.ConnectToDb()
	defer cancel()
	engine := gin.Default()
	api := engine.Group("/api")
	handlers.RoutHandler(api)
	PORT := os.Getenv("PORT")
	engine.Run(":" + PORT)
}
