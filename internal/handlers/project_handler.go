package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"logarithm/internal/dto"
	"logarithm/internal/models"
	"logarithm/internal/services/interfaces"
	"net/http"
)

type ProjectHandler struct {
	routerGroup *gin.RouterGroup
	service     interfaces.IStudentService
}

func (handler *ProjectHandler) RegisterProjectsGroup() {
	group := handler.routerGroup.Group("/students")

	group.GET("", handler.getAll)
	group.GET("/:id", handler.get)
	group.POST("", handler.create)
	group.PUT("/:id", handler.update)
	group.DELETE("/:id", handler.delete)
}

func (handler *ProjectHandler) getAll(context *gin.Context) {
	students, err := handler.service.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.NewErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	context.JSON(200, dto.NewSuccessResponse(students))
}

func (handler *ProjectHandler) get(context *gin.Context) {
	studentId := context.Param("id")

	student, err := handler.service.Get(studentId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	context.JSON(200, dto.NewSuccessResponse(student))
}

func (handler *ProjectHandler) create(context *gin.Context) {
	student := models.Student{}

	// Request body to struct binding
	err := context.ShouldBindJSON(&student)
	if err != nil {
		context.JSON(400, dto.NewErrorResponse(400, err.Error()))
		return
	}

	// Project insertion
	insertedStudent, err := handler.service.Create(student)
	if err != nil {
		context.JSON(500, dto.NewErrorResponse(500, err.Error()))
		return
	}

	context.JSON(201, dto.NewSuccessResponse(insertedStudent))
}

func (handler *ProjectHandler) update(context *gin.Context) {
	student := dto.StudentUpdateDTO{}
	studentId := context.Param("id")

	err := context.ShouldBindJSON(&student)
	if err != nil {
		context.JSON(400, dto.NewErrorResponse(400, err.Error()))
		return
	}

	updatedProject, err := handler.service.Update(studentId, student)
	var pqErr *pq.Error
	if errors.Is(err, sql.ErrNoRows) || (errors.As(err, &pqErr) && pqErr.Code == "22P02") {
		context.JSON(404, dto.NewErrorResponse(404, "Student not found"))
		return
	}
	if err != nil {
		context.JSON(500, dto.NewErrorResponse(500, err.Error()))
		return
	}

	context.JSON(200, dto.NewSuccessResponse(updatedProject))
}

func (handler *ProjectHandler) delete(context *gin.Context) {
	studentId := context.Param("id")

	deletedProject, err := handler.service.Delete(studentId)
	if err != nil {
		context.JSON(404, dto.NewErrorResponse(404, "Student not found"))
		return
	}

	context.JSON(200, dto.NewSuccessResponse(deletedProject))
}
