package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"logarithm/internal/handlers"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	handlers.RegisterProjectHandlers(v1)
	router.Run()
}
