package user

type CreateUserResponse struct {
	Message string `json:"message"`
}

type UpdateUserInput struct {
	Email *string `json:"email" validate:"omitempty,email"`
	Password *string `json:"password" validate:"omitempty,min=8,max=32"`
}

type BadResponse struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Fields map[string]string `json:"fields"`
}
