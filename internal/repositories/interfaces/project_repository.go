package interfaces

import "logarithm/internal/models"

type IProjectRepository interface {
	GetAll() []models.Project
}
