package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ilyasssan2/golangGin/handlers"
	"github.com/ilyasssan2/golangGin/utils"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cancel := utils.ConnectToDb()
	defer cancel()
	engine := gin.Default()
	engine.Static("/static", "./public")
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("CLIEN_URL")},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
	}))
	api := engine.Group("/api")
	handlers.RoutHandler(api)
	PORT := os.Getenv("PORT")
	engine.Run(":" + PORT)
}
