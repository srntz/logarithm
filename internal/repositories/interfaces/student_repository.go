package interfaces

import (
	"logarithm/internal/dto"
	"logarithm/internal/models"
)

type IStudentRepository interface {
	GetAll() []models.Student
	Create(student models.Student) (models.Student, error)
	Update(studentId string, student dto.StudentUpdateDTO) (models.Student, error)
	Delete(studentId string) (models.Student, error)
}
