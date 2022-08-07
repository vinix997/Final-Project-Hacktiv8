package service

import (
	"context"
	"errors"

	"github.com/vinix997/Final-Project-Hacktiv8/entity"
)

type PhotoService interface {
	CreatePhoto(ctx context.Context, photo *entity.Photo) (*entity.Photo, *entity.Error)
	UpdatePhoto(ctx context.Context, photo *entity.Photo, id int) (*entity.Photo, *entity.Error)
	GetPhoto(ctx context.Context) ([]entity.PhotoResponse, *entity.Error)
	DeletePhoto(ctx context.Context, id int) (string, *entity.Error)
}

type PhotoRepository interface {
	CreatePhoto(ctx context.Context, photo *entity.Photo) (*entity.Photo, *entity.Error)
	UpdatePhoto(ctx context.Context, photo *entity.Photo, id int) (*entity.Photo, *entity.Error)
	GetPhoto(ctx context.Context) ([]entity.PhotoResponse, *entity.Error)
	DeletePhoto(ctx context.Context, id int) (string, *entity.Error)
}

type PhotoSvc struct {
	photoRepo PhotoRepository
}

func NewPhotoSvc(photoRepo PhotoRepository) PhotoService {
	return &PhotoSvc{
		photoRepo: photoRepo,
	}
}

func (u *PhotoSvc) CreatePhoto(ctx context.Context, photo *entity.Photo) (*entity.Photo, *entity.Error) {
	err := validatePhoto(photo)
	if err != nil {
		return nil, &entity.Error{IsError: true, Message: err.Error()}
	}
	return u.photoRepo.CreatePhoto(ctx, photo)
}
func (u *PhotoSvc) UpdatePhoto(ctx context.Context, photo *entity.Photo, id int) (*entity.Photo, *entity.Error) {
	err := validatePhoto(photo)
	if err != nil {
		return nil, &entity.Error{IsError: true, Message: err.Error()}
	}
	return u.photoRepo.UpdatePhoto(ctx, photo, id)
}
func (u *PhotoSvc) GetPhoto(ctx context.Context) ([]entity.PhotoResponse, *entity.Error) {
	return u.photoRepo.GetPhoto(ctx)
}
func (u *PhotoSvc) DeletePhoto(ctx context.Context, id int) (string, *entity.Error) {
	return u.photoRepo.DeletePhoto(ctx, id)
}

func validatePhoto(photo *entity.Photo) error {
	if photo.Title == "" {
		return errors.New("Title cannot be empty")
	}
	if photo.PhotoUrl == "" {
		return errors.New("Photo url cannot be empty")
	}
	return nil
}
