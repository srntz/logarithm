package interfaces

import (
	"logarithm/internal/dto"
	"logarithm/internal/models"
)

type IProjectRepository interface {
	GetAll() []models.Project
	Create(project dto.ProjectInsertDTO) (models.Project, error)
	Update(projectId string, project dto.ProjectUpdateDTO) (models.Project, error)
}
