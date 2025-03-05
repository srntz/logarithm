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

func (repository *projectRepository) Create(project models.ProjectInsertDTO) (models.Project, error) {
	insertedProject := models.Project{}

	query := `INSERT INTO project (name, allowed_origin) VALUES ($1, $2) RETURNING id, name, allowed_origin;`

	err := repository.db.QueryRow(query, project.Name, project.AllowedOrigin).Scan(&insertedProject.Id, &insertedProject.Name, &insertedProject.AllowedOrigin)
	if err != nil {
		return models.Project{}, err
	}

	return insertedProject, nil
}

func (repository *projectRepository) Update(projectId string, project models.ProjectUpdateDTO) (models.Project, error) {
	updatedProject := models.Project{}

	query := `UPDATE project SET 
                   name = COALESCE($1, name),
				   allowed_origin = COALESCE($2, allowed_origin)
				WHERE id = $3
				RETURNING id, name, allowed_origin;`

	err := repository.db.QueryRow(query, project.Name, project.AllowedOrigin, projectId).Scan(&updatedProject.Id, &updatedProject.Name, &updatedProject.AllowedOrigin)
	if err != nil {
		return models.Project{}, err
	}

	return updatedProject, nil

}
