package user

import (
	passwordhashing "github.com/fathnasrudin/todogo/internal/common/passwordHashing"
	"github.com/google/uuid"
)

type IUserService interface {
	List() ([]GetUserResponse, error)
	Create(d CreateUserInput) (error)
	Update(id string, d UpdateUserInput) (error)
	Delete(id string) (error)
}

type UserService struct {
	repo UserRepo
}

func NewUserService(r UserRepo) *UserService{
	return &UserService{
		repo: r,
	}
}

type CreateUserInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func newUser(uData CreateUserInput) (User, error){
	IDByte, err := uuid.NewV7();
	if err != nil {return User{}, err}

	// hash password
	hashedPassword, err := passwordhashing.HashPassword(uData.Password)
	if err != nil {
		return User{}, err
	}

	return User{
		ID: IDByte.String(),
		Email: uData.Email,
		Password: hashedPassword,
	}, nil
}

func (s UserService) List() ([]GetUserResponse, error) {
	users, err := s.repo.List()
	if err != nil {return nil, err}
	return users, nil
}

func (s UserService) Create(uData CreateUserInput) (error) {
	user, err := newUser(uData)
	if err != nil {
		return err
	}

	if err := s.repo.Create(user); err != nil {return err}
	return nil
}

func (s UserService) Update(id string, d UpdateUserInput) (error) {

	if (d.Password != nil) {
		hashed, err := passwordhashing.HashPassword(*d.Password)
		if err != nil {return err}
		d.Password = &hashed
	}

	err := s.repo.Update(id, d)
	return err
}


func (s UserService) Delete(id string) (error) {
	return s.repo.Delete(id);
}