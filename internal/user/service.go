package user

type IUserService interface {
	List() ([]GetUserResponse, error)
}

type UserService struct {
	repo UserRepo
}

func NewUserService(r UserRepo) *UserService{
	return &UserService{
		repo: r,
	}
}

func (s UserService) List() ([]GetUserResponse, error) {
	users, err := s.repo.List()
	if err != nil {return nil, err}
	return users, nil
}