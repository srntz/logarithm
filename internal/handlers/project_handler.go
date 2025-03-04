package handlers

import (
	"github.com/gin-gonic/gin"
	"logarithm/internal/services/interfaces"
)

type ProjectHandler struct {
	routerGroup *gin.RouterGroup
	service     interfaces.IProjectService
}

func (handler *ProjectHandler) RegisterProjectsGroup() {
	group := handler.routerGroup.Group("/projects")

	group.GET("", handler.getAll)
}

func (handler *ProjectHandler) getAll(context *gin.Context) {
	context.JSON(200, handler.service.GetAll(context))
}
