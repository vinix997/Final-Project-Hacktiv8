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

func TestNewCommentSvc(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("Test Initiate New Photo Service", func(t *testing.T) {
		mockCommentRepo := mock_service.NewMockCommentRepository(ctrl)
		commentService := service.NewCommentSvc(mockCommentRepo)
		require.NotNil(t, commentService)
	})
}

func TestCommentService(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Empty comment", func(t *testing.T) {
		mockCommentRepo := mock_service.NewMockCommentRepository(ctrl)
		commentService := service.NewCommentSvc(mockCommentRepo)
		res, err := commentService.CreateComment(context.Background(), &entity.Comment{
			Message: "",
		})
		require.Error(t, errors.New(err.Message))
		require.Equal(t, &entity.Error{IsError: true, Message: "Message cannot be empty"}, err)
		require.Nil(t, res)
	})
	t.Run("Success create", func(t *testing.T) {
		mockCommentRepo := mock_service.NewMockCommentRepository(ctrl)
		commentService := service.NewCommentSvc(mockCommentRepo)
		comment := &entity.Comment{
			Message: "Karen",
		}
		commentReq := &entity.Comment{
			Message: "Karen",
		}
		mockCommentRepo.EXPECT().CreateComment(gomock.Any(), gomock.Any()).Return(commentReq, nil)
		res, err := commentService.CreateComment(context.Background(), comment)
		require.NotNil(t, res)
		require.Nil(t, err)
		require.Equal(t, comment.Message, res.Message)
	})
}
