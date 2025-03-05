package interfaces

import "logarithm/internal/models"

type IProjectRepository interface {
	GetAll() []models.Project
	Create(project models.ProjectInsertDTO) (models.Project, error)
	Update(projectId string, project models.ProjectUpdateDTO) (models.Project, error)
}
