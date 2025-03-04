package main

import (
	"database/sql"
	"fmt"
	"log"
	"logarithm/db"
	"logarithm/internal/models"
)

var projects = []models.ProjectInsertDTO{
	{"Project 1", "http://example.com"},
}

var errors = []models.ErrorInsertDTO{
	{"", "InternalServerError", "Error: Index out of bounds", "stack_trace", "http://localhost:9999", ""},
	{"", "InternalServerError", "Error: Integer overflow", "stack_trace", "http://localhost:9999", ""},
}

var dbInstance *sql.DB

func main() {
	dbInstance = db.Connect()
	dropAll()
	loadSchema()
	for i := 0; i < len(projects); i++ {
		loadProjects(&projects[i], errors)
	}
}

func dropAll() {
	query := `DROP TABLE IF EXISTS project, error CASCADE`
	_, err := dbInstance.Exec(query)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Tables dropped.")
	}
}

func loadSchema() {
	query := `	CREATE TABLE IF NOT EXISTS project (
					id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
					name varchar(64) NOT NULL,
					allowed_origin varchar(255) NOT NULL);

				CREATE TABLE IF NOT EXISTS error (
					id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
					project_id UUID NOT NULL,
					type TEXT NOT NULL,
					message TEXT NOT NULL,
					stack_trace TEXT NOT NULL,
					logged_at TIMESTAMP NOT NULL default now(),
					request_url TEXT NOT NULL,
					metadata TEXT NOT NULL);`

	_, err := dbInstance.Exec(query)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Tables created.")
	}
}

func loadProjects(project *models.ProjectInsertDTO, errors []models.ErrorInsertDTO) {
	projectQuery := `INSERT INTO project (name, allowed_origin) VALUES ($1, $2) RETURNING id;`
	errorQuery := `INSERT INTO error (project_id, type, message, stack_trace, request_url, metadata) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	var projectId string
	err := dbInstance.QueryRow(projectQuery, project.Name, project.AllowedOrigin).Scan(&projectId)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Project %s inserted \n", projectId)
	}

	for i := 0; i < len(errors); i++ {
		errors[i].ProjectId = projectId
		var errorId string
		err := dbInstance.QueryRow(errorQuery, errors[i].ProjectId, errors[i].Type, errors[i].Message, errors[i].StackTrace, errors[i].RequestURL, errors[i].Metadata).Scan(&errorId)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Error %s inserted \n", errorId)
		}
	}
}
