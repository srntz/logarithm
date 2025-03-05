package models

type Project struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	AllowedOrigin string `json:"allowed_origin"`
}
