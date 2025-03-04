package handlers

import "github.com/gin-gonic/gin"

func RegisterProjectHandlers(r *gin.RouterGroup) {
	r.GET("/projects", getProjects)
}

func getProjects(c *gin.Context) {
	c.JSON(200, "projects")
}
