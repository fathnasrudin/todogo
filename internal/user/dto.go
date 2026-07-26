package user

type CreateUserResponse struct {
	Message string `json:"message"`
}

type UpdateUserInput struct {
	Email *string `json:"email"`
	Password *string `json:"password"`
}
