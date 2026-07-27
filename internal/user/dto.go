package user

type CreateUserResponse struct {
	Message string `json:"message"`
}

type UpdateUserInput struct {
	Email *string `json:"email"`
	Password *string `json:"password"`
}

type BadResponse struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Fields map[string]string `json:"fields"`
}
