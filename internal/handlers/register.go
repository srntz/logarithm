package handlers

import (
	"github.com/gin-gonic/gin"
	repositories "logarithm/internal/repositories/implementations"
	services "logarithm/internal/services/implementations"
)

func Register(r *gin.RouterGroup) {
	projectsHandler := ProjectHandler{routerGroup: r, service: services.NewProjectService(repositories.NewProjectRepository())}
	projectsHandler.RegisterProjectsGroup()
}
