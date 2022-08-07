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

func TestNewPhotoSvc(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("Test Initiate New Photo Service", func(t *testing.T) {
		mockPhotoRepo := mock_service.NewMockSocialRepository(ctrl)
		photoService := service.NewSocialSvc(mockPhotoRepo)
		require.NotNil(t, photoService)
	})
}

func TestPhotoService(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty title", func(t *testing.T) {
		mockPhotoRepo := mock_service.NewMockPhotoRepository(ctrl)
		photoService := service.NewPhotoSvc(mockPhotoRepo)
		res, err := photoService.CreatePhoto(context.Background(), &entity.Photo{
			Title:    "",
			PhotoUrl: "",
		})
		require.Error(t, errors.New(err.Message))
		require.Equal(t, &entity.Error{IsError: true, Message: "Title cannot be empty"}, err)
		require.Nil(t, res)
	})
	t.Run("Empty photo url", func(t *testing.T) {
		mockPhotoRepo := mock_service.NewMockPhotoRepository(ctrl)
		photoService := service.NewPhotoSvc(mockPhotoRepo)
		res, err := photoService.CreatePhoto(context.Background(), &entity.Photo{
			Title:    "Title test",
			PhotoUrl: "",
		})
		require.Error(t, errors.New(err.Message))
		require.Equal(t, &entity.Error{IsError: true, Message: "Photo url cannot be empty"}, err)
		require.Nil(t, res)
	})
}
