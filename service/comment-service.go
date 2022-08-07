package service

import (
	"context"
	"errors"

	"github.com/vinix997/Final-Project-Hacktiv8/entity"
)

type CommentService interface {
	CreateComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, *entity.Error)
	UpdateComment(ctx context.Context, comment *entity.Comment, id int) (*entity.CommentUpdateResponse, *entity.Error)
	GetComment(ctx context.Context) ([]entity.CommentResponse, *entity.Error)
	DeleteComment(ctx context.Context, id int) (string, *entity.Error)
}

type CommentRepository interface {
	CreateComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, *entity.Error)
	UpdateComment(ctx context.Context, comment *entity.Comment, id int) (*entity.CommentUpdateResponse, *entity.Error)
	GetComment(ctx context.Context) ([]entity.CommentResponse, *entity.Error)
	DeleteComment(ctx context.Context, id int) (string, *entity.Error)
}

type CommentSvc struct {
	commentRepo CommentRepository
}

func NewCommentSvc(commentRepo CommentRepository) CommentService {
	return &CommentSvc{
		commentRepo: commentRepo,
	}
}

func (u *CommentSvc) CreateComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, *entity.Error) {
	err := validateComment(comment)
	if err != nil {
		return nil, &entity.Error{IsError: true, Message: err.Error()}
	}
	return u.commentRepo.CreateComment(ctx, comment)
}
func (u *CommentSvc) UpdateComment(ctx context.Context, comment *entity.Comment, id int) (*entity.CommentUpdateResponse, *entity.Error) {
	return u.commentRepo.UpdateComment(ctx, comment, id)
}
func (u *CommentSvc) GetComment(ctx context.Context) ([]entity.CommentResponse, *entity.Error) {
	return u.commentRepo.GetComment(ctx)
}
func (u *CommentSvc) DeleteComment(ctx context.Context, id int) (string, *entity.Error) {
	return u.commentRepo.DeleteComment(ctx, id)
}

func validateComment(comment *entity.Comment) error {
	if comment.Message == "" {
		return errors.New("Message cannot be empty")
	}
	return nil
}
