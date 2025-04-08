package models

type Student struct {
	StudentId   string `json:"studentId"`
	StudentName string `json:"studentname"`
	CourseName  string `json:"courseName"`
	Date        string `json:"date"`
}
