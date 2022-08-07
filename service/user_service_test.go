package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vinix997/Final-Project-Hacktiv8/entity"
	mock_service "github.com/vinix997/Final-Project-Hacktiv8/mock/service"
	"github.com/vinix997/Final-Project-Hacktiv8/service"
)

func TestNewUserSvc(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("Test Initiate New User Service", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		require.NotNil(t, userService)
	})
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty username", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		res, err := userService.CreateUser(context.Background(), &entity.User{
			Username: "",
			Email:    "test@gmail.com",
			Password: "dummypassword",
		})
		require.Error(t, err)
		require.Equal(t, errors.New("Username cannot be empty"), err)
		require.Nil(t, res)
	})
	t.Run("Invalid email", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		res, err := userService.UpdateUser(context.Background(), &entity.User{
			Username: "",
			Email:    "tesadadsa",
			Password: "dummypassword",
		}, 1)
		require.Error(t, errors.New(err.Message))
		require.Equal(t, &entity.Error{IsError: true, Message: "Invalid email"}, err)
		require.Nil(t, res)
	})
	t.Run("Success insert", func(t *testing.T) {
		mockUserRepo := mock_service.NewMockUserRepository(ctrl)
		userService := service.NewUserSvc(mockUserRepo)
		user := &entity.User{
			Username: "testuser",
			Email:    "test@gmail.com",
			Password: "dummypassword",
			Age:      19,
		}
		response := &entity.User{
			Username: "testuser",
			Email:    "test@gmail.com",
			Password: "dummypassword",
			Age:      19,
		}
		mockUserRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(response, nil)
		res, err := userService.CreateUser(context.Background(), user)
		require.Nil(t, err)
		require.NotNil(t, res)
		require.Equal(t, user.Username, res.Username)
	})
}
