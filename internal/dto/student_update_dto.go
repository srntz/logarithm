package dto

type StudentUpdateDTO struct {
	StudentName *string `json:"studentName"`
	CourseName  *string `json:"courseName"`
	Date        *string `json:"date"`
}
