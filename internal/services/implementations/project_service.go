package implementations

import (
	"logarithm/internal/dto"
	"logarithm/internal/models"
	"logarithm/internal/repositories/interfaces"
)

type projectService struct {
	repository interfaces.IProjectRepository
}

func NewProjectService(repository interfaces.IProjectRepository) *projectService {
	return &projectService{repository: repository}
}

func (s *projectService) GetAll() []models.Project {
	return s.repository.GetAll()
}

func (s *projectService) Create(project dto.ProjectInsertDTO) (models.Project, error) {
	return s.repository.Create(project)
}

func (s *projectService) Update(projectId string, project dto.ProjectUpdateDTO) (models.Project, error) {
	return s.repository.Update(projectId, project)
}
