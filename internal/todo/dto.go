package todo

type ResponseMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type BadResponse struct {
	Message string `json:"message"`
}

type UpdateTaskInput struct {
	Title *string `json:"title" validate:"omitempty,min=1"`
	IsDone *bool `json:"is_done" validate:"omitempty"`
}

type UpdateTaskResponse struct {
	Message string `json:"message"`
}

type CreateTaskInput struct {
	Title string `json:"title" validate:"required,min=1"`
}

type CreateTaskResponse struct {
	Message string `json:"message"`
}
