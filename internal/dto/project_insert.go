package dto

type ProjectInsertDTO struct {
	Name          string `json:"name" binding:"required"`
	AllowedOrigin string `json:"allowed_origin" binding:"required"`
}
