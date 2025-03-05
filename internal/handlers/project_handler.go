package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"logarithm/internal/dto/responses"
	"logarithm/internal/models"
	"logarithm/internal/services/interfaces"
)

type ProjectHandler struct {
	routerGroup *gin.RouterGroup
	service     interfaces.IProjectService
}

func (handler *ProjectHandler) RegisterProjectsGroup() {
	group := handler.routerGroup.Group("/projects")

	group.GET("", handler.getAll)
	group.POST("", handler.create)
	group.PUT("/:id", handler.update)
}

func (handler *ProjectHandler) getAll(context *gin.Context) {
	context.JSON(200, handler.service.GetAll())
}

func (handler *ProjectHandler) create(context *gin.Context) {
	project := models.ProjectInsertDTO{}

	// Request body to struct binding
	err := context.ShouldBindJSON(&project)
	if err != nil {
		context.JSON(400, responses.NewErrorResponse(400, err.Error()))
		return
	}

	// Project insertion
	insertedProject, err := handler.service.Create(project)
	if err != nil {
		context.JSON(500, responses.NewErrorResponse(500, err.Error()))
		return
	}

	context.JSON(200, responses.NewSuccessResponse(insertedProject))
}

func (handler *ProjectHandler) update(context *gin.Context) {
	project := models.ProjectUpdateDTO{}
	projectId := context.Param("id")

	err := context.ShouldBindJSON(&project)
	if err != nil {
		context.JSON(400, responses.NewErrorResponse(400, err.Error()))
		return
	}

	updatedProject, err := handler.service.Update(projectId, project)
	var pqErr *pq.Error
	if errors.Is(err, sql.ErrNoRows) || (errors.As(err, &pqErr) && pqErr.Code == "22P02") {
		context.JSON(404, responses.NewErrorResponse(404, "Project not found"))
		return
	}
	if err != nil {
		context.JSON(500, responses.NewErrorResponse(500, err.Error()))
		return
	}

	context.JSON(200, responses.NewSuccessResponse(updatedProject))
}
