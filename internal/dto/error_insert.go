package dto

type ErrorInsertDTO struct {
	ProjectId  string `json:"project_id"`
	Type       string `json:"type"`
	Message    string `json:"message"`
	StackTrace string `json:"stack_trace"`
	RequestURL string `json:"request_url"`
	Metadata   string `json:"metadata"`
}
