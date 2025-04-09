package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"logarithm/internal/handlers"
	"net/http"
)

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("./internal/templates/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1 := router.Group("/api/v1")
	handlers.Register(v1)

	router.Run()
}
