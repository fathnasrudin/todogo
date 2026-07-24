package todo

type ResponseMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type BadResponse struct {
	Message string `json:"message"`
}

type UpdateTaskInput struct {
	Title *string `json:"title"`
	IsDone *bool `json:"is_done"`
}

type UpdateTaskResponse struct {
	Message string `json:"message"`
}

type CreateTaskInput struct {
	Title string `json:"title"`
}

type CreateTaskResponse struct {
	Message string `json:"message"`
}
