package implementations

import (
	"logarithm/internal/dto"
	"logarithm/internal/models"
	"logarithm/internal/repositories/interfaces"
)

type studentService struct {
	repository interfaces.IStudentRepository
}

func NewStudentService(repository interfaces.IStudentRepository) *studentService {
	return &studentService{repository: repository}
}

func (s *studentService) Get(studentId string) (*models.Student, error) {
	return s.repository.Get(studentId)
}

func (s *studentService) GetAll() ([]models.Student, error) {
	return s.repository.GetAll()
}

func (s *studentService) Create(student models.Student) (models.Student, error) {
	return s.repository.Create(student)
}

func (s *studentService) Update(studentId string, student dto.StudentUpdateDTO) (models.Student, error) {
	return s.repository.Update(studentId, student)
}

func (s *studentService) Delete(projectId string) (models.Student, error) {
	return s.repository.Delete(projectId)
}
