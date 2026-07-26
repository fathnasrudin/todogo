package user

type User struct {
	ID string
	Email string
	Password string
}

type GetUserResponse struct {
	ID string
	Email string
}