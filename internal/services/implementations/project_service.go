package implementations

import (
	"github.com/gin-gonic/gin"
	"logarithm/internal/models"
	"logarithm/internal/repositories/interfaces"
)

type projectService struct {
	repository interfaces.IProjectRepository
}

func NewProjectService(repository interfaces.IProjectRepository) *projectService {
	return &projectService{repository: repository}
}

func (s *projectService) GetAll(c *gin.Context) []models.Project {
	return s.repository.GetAll()
}
