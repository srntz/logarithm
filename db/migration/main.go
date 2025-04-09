package main

import (
	"database/sql"
	"fmt"
	"log"
	"logarithm/db"
	"logarithm/internal/models"
)

var students = []models.Student{
	{StudentId: "1", StudentName: "Student 1", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "2", StudentName: "Student 2", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "3", StudentName: "Student 3", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "4", StudentName: "Student 4", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "5", StudentName: "Student 5", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "6", StudentName: "Student 6", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "7", StudentName: "Student 7", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "8", StudentName: "Student 8", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "9", StudentName: "Student 9", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "10", StudentName: "Student 10", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "11", StudentName: "Student 11", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "12", StudentName: "Student 12", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "13", StudentName: "Student 13", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "14", StudentName: "Student 14", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "15", StudentName: "Student 15", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "16", StudentName: "Student 16", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "17", StudentName: "Student 17", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "18", StudentName: "Student 18", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "19", StudentName: "Student 19", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "20", StudentName: "Student 20", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "21", StudentName: "Student 21", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "22", StudentName: "Student 22", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "23", StudentName: "Student 23", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "24", StudentName: "Student 24", CourseName: "OS", Date: "01/01/1970"},
	{StudentId: "25", StudentName: "Student 25", CourseName: "OS", Date: "01/01/1970"},
}

var dbInstance *sql.DB

func main() {
	dbInstance = db.Connect()
	dropAll()
	loadSchema()
	for _, student := range students {
		insertStudent(&student)
	}
	log.Println("Data Inserted.")
}

func dropAll() {
	query := `DROP TABLE IF EXISTS student CASCADE`
	_, err := dbInstance.Exec(query)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Tables dropped.")
	}
}

func loadSchema() {
	query := `CREATE TABLE IF NOT EXISTS student (
					student_id VARCHAR(64) PRIMARY KEY,
	    			student_name TEXT,
	    			course_name TEXT,
	    			date VARCHAR(10));`

	_, err := dbInstance.Exec(query)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Tables created.")
	}
}

func insertStudent(student *models.Student) {
	query := `INSERT INTO student (student_id, student_name, course_name, date) VALUES ($1, $2, $3, $4);`

	_, err := dbInstance.Exec(query, student.StudentId, student.StudentName, student.CourseName, student.Date)
	if err != nil {
		log.Fatal(err)
	}
}
