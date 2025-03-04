package models

type Error struct {
	Id         string `json:"id"`
	ProjectId  string `json:"project_id"`
	Type       string `json:"type"`
	Message    string `json:"message"`
	StackTrace string `json:"stack_trace"`
	LoggedAt   int64  `json:"logged_at"`
	RequestURL string `json:"request_url"`
	Metadata   string `json:"metadata"`
}

type ErrorInsertDTO struct {
	ProjectId  string `json:"project_id"`
	Type       string `json:"type"`
	Message    string `json:"message"`
	StackTrace string `json:"stack_trace"`
	RequestURL string `json:"request_url"`
	Metadata   string `json:"metadata"`
}
