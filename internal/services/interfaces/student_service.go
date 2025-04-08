package interfaces

import (
	"logarithm/internal/dto"
	"logarithm/internal/models"
)

type IStudentService interface {
	Get(studentId string) (*models.Student, error)
	GetAll() ([]models.Student, error)
	Create(student models.Student) (models.Student, error)
	Update(studentId string, student dto.StudentUpdateDTO) (models.Student, error)
	Delete(studentId string) (models.Student, error)
}
