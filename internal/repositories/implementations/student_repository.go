package implementations

import (
	"database/sql"
	"log"
	"logarithm/db"
	"logarithm/internal/dto"
	"logarithm/internal/models"
)

type studentRepository struct {
	db *sql.DB
}

func NewStudentRepository() *studentRepository {
	return &studentRepository{db: db.Connect()}
}

func (repository *studentRepository) GetAll() []models.Student {
	students := []models.Student{}
	query := `SELECT student_id, student_name, course_name, date FROM student;`

	rows, err := repository.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		student := models.Student{}

		err := rows.Scan(&student.StudentId, &student.StudentName, &student.CourseName, &student.Date)
		if err != nil {
			log.Fatal(err)
		}

		students = append(students, student)
	}

	return students
}

func (repository *studentRepository) Create(student models.Student) (models.Student, error) {
	insertedStudent := models.Student{}

	query := `INSERT INTO student (student_id, student_name, course_name, date) VALUES ($1, $2, $3, $4) RETURNING student_id, student_name, course_name, date;`

	err := repository.db.QueryRow(query, student.StudentId, student.StudentName, student.CourseName, student.Date).Scan(&insertedStudent.StudentId, &insertedStudent.StudentName, &insertedStudent.CourseName, &insertedStudent.Date)
	if err != nil {
		return models.Student{}, err
	}

	return insertedStudent, nil
}

func (repository *studentRepository) Update(studentId string, student dto.StudentUpdateDTO) (models.Student, error) {
	updatedStudent := models.Student{}

	query := `UPDATE student SET 
                   student_name = COALESCE($1, student_name),
				   course_name = COALESCE($2, course_name),
				   date = COALESCE($3, date)
				WHERE student_id = $4
				RETURNING student_id, student_name, course_name, date;`

	err := repository.db.QueryRow(query, student.StudentName, student.CourseName, student.Date, studentId).Scan(&updatedStudent.StudentId, &updatedStudent.StudentName, &updatedStudent.CourseName, &updatedStudent.Date)
	if err != nil {
		return models.Student{}, err
	}

	return updatedStudent, nil

}

func (repository *studentRepository) Delete(studentId string) (models.Student, error) {
	deletedStudent := models.Student{}

	query := `DELETE FROM student WHERE student_id = $1 RETURNING student_id, student_name, course_name, date;`

	err := repository.db.QueryRow(query, studentId).Scan(&deletedStudent.StudentId, &deletedStudent.StudentName, &deletedStudent.CourseName, &deletedStudent.Date)
	if err != nil {
		return models.Student{}, err
	}

	return deletedStudent, nil
}
