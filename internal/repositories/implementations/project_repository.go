package implementations

import (
	"database/sql"
	"log"
	"logarithm/db"
	"logarithm/internal/models"
)

type projectRepository struct {
	db *sql.DB
}

func NewProjectRepository() *projectRepository {
	return &projectRepository{db: db.Connect()}
}

func (repository *projectRepository) GetAll() []models.Project {
	projects := []models.Project{}
	query := `SELECT id, name, allowed_origin FROM project;`

	rows, err := repository.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		project := models.Project{}

		err := rows.Scan(&project.Id, &project.Name, &project.AllowedOrigin)
		if err != nil {
			log.Fatal(err)
		}

		projects = append(projects, project)
	}

	return projects
}
