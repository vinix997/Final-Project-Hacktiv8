package service

import (
	"context"
	"errors"
	"net/mail"

	"github.com/vinix997/Final-Project-Hacktiv8/entity"
)

type UserService interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User, id int) (*entity.User, *entity.Error)
	Login(ctx context.Context, user *entity.User) (string, *entity.Error)
	DeleteUser(ctx context.Context, id int) (string, *entity.Error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User, id int) (*entity.User, *entity.Error)
	Login(ctx context.Context, user *entity.User) (string, *entity.Error)
	DeleteUser(ctx context.Context, id int) (string, *entity.Error)
}

type UserSvc struct {
	userRepo UserRepository
}

func NewUserSvc(userRepo UserRepository) UserService {
	return &UserSvc{
		userRepo: userRepo,
	}
}

func (u *UserSvc) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := validateUser(user)
	if err != nil {
		return nil, err
	}
	return u.userRepo.CreateUser(ctx, user)
}
func (u *UserSvc) UpdateUser(ctx context.Context, user *entity.User, id int) (*entity.User, *entity.Error) {
	err := validateUser(user)
	if err != nil {
		return nil, &entity.Error{
			IsError: true,
			Message: err.Error(),
		}
	}
	return u.userRepo.UpdateUser(ctx, user, id)
}
func (u *UserSvc) Login(ctx context.Context, user *entity.User) (string, *entity.Error) {
	return u.userRepo.Login(ctx, user)
}
func (u *UserSvc) DeleteUser(ctx context.Context, id int) (string, *entity.Error) {
	return u.userRepo.DeleteUser(ctx, id)
}

func validateUser(user *entity.User) error {
	if user == nil {
		return errors.New("User cannot be empty")
	}
	if valid(user.Email) == false {
		return errors.New("Invalid email")
	}
	if user.Username == "" {
		return errors.New("Username cannot be empty")
	}
	if len(user.Password) < 6 {
		return errors.New("Password must be at least 6 characters")
	}
	if user.Age < 8 {
		return errors.New("You are too young")
	}

	return nil
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
