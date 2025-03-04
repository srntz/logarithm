package interfaces

import (
	"github.com/gin-gonic/gin"
	"logarithm/internal/models"
)

type IProjectService interface {
	GetAll(c *gin.Context) []models.Project
}
