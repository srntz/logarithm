package models

type Project struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	AllowedOrigin string `json:"allowed_origin"`
}

type ProjectInsertDTO struct {
	Name          string `json:"name" binding:"required"`
	AllowedOrigin string `json:"allowed_origin" binding:"required"`
}

type ProjectUpdateDTO struct {
	Name          *string `json:"name"`
	AllowedOrigin *string `json:"allowed_origin"`
}
