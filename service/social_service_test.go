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

func TestNewSocialSvc(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("Test Initiate New Social Service", func(t *testing.T) {
		mockSocialRepo := mock_service.NewMockSocialRepository(ctrl)
		socialService := service.NewSocialSvc(mockSocialRepo)
		require.NotNil(t, socialService)
	})
}

func TestSocialService(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty name", func(t *testing.T) {
		mockSocialRepo := mock_service.NewMockSocialRepository(ctrl)
		socialService := service.NewSocialSvc(mockSocialRepo)
		res, err := socialService.CreateSocial(context.Background(), &entity.SocialMedia{
			Name:           "",
			SocialMediaUrl: "",
		})
		require.Error(t, errors.New(err.Message))
		require.Equal(t, &entity.Error{IsError: true, Message: "Name cannot be empty"}, err)
		require.Nil(t, res)
	})
	t.Run("Empty url", func(t *testing.T) {
		mockSocialRepo := mock_service.NewMockSocialRepository(ctrl)
		socialService := service.NewSocialSvc(mockSocialRepo)
		res, err := socialService.CreateSocial(context.Background(), &entity.SocialMedia{
			Name:           "asdasd",
			SocialMediaUrl: "",
		})
		require.Error(t, errors.New(err.Message))
		require.Equal(t, &entity.Error{IsError: true, Message: "Social media url cannot be empty"}, err)
		require.Nil(t, res)
	})
}
