package dto

type ProjectUpdateDTO struct {
	Name          *string `json:"name"`
	AllowedOrigin *string `json:"allowed_origin"`
}
