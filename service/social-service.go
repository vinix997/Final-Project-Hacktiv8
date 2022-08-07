package service

import (
	"context"
	"errors"

	"github.com/vinix997/Final-Project-Hacktiv8/entity"
)

type SocialService interface {
	CreateSocial(ctx context.Context, Social *entity.SocialMedia) (*entity.SocialMedia, *entity.Error)
	UpdateSocial(ctx context.Context, Social *entity.SocialMedia, id int) (*entity.SocialMedia, *entity.Error)
	DeleteSocial(ctx context.Context, id int) (string, *entity.Error)
	GetSocial(ctx context.Context) (*entity.CustomSocialMediaResponse, *entity.Error)
}

type SocialRepository interface {
	CreateSocial(ctx context.Context, Social *entity.SocialMedia) (*entity.SocialMedia, *entity.Error)
	UpdateSocial(ctx context.Context, Social *entity.SocialMedia, id int) (*entity.SocialMedia, *entity.Error)
	DeleteSocial(ctx context.Context, id int) (string, *entity.Error)
	GetSocial(ctx context.Context) (*entity.CustomSocialMediaResponse, *entity.Error)
}

type SocialSvc struct {
	socialRepo SocialRepository
}

func NewSocialSvc(socialRepo SocialRepository) SocialService {
	return &SocialSvc{
		socialRepo: socialRepo,
	}
}

func (u *SocialSvc) CreateSocial(ctx context.Context, Social *entity.SocialMedia) (*entity.SocialMedia, *entity.Error) {
	err := validateSocial(Social)
	if err != nil {
		return nil, &entity.Error{IsError: true, Message: err.Error()}
	}
	return u.socialRepo.CreateSocial(ctx, Social)
}
func (u *SocialSvc) UpdateSocial(ctx context.Context, Social *entity.SocialMedia, id int) (*entity.SocialMedia, *entity.Error) {
	err := validateSocial(Social)
	if err != nil {
		return nil, &entity.Error{IsError: true, Message: err.Error()}
	}
	return u.socialRepo.UpdateSocial(ctx, Social, id)
}

func (u *SocialSvc) DeleteSocial(ctx context.Context, id int) (string, *entity.Error) {
	return u.socialRepo.DeleteSocial(ctx, id)
}

func (u *SocialSvc) GetSocial(ctx context.Context) (*entity.CustomSocialMediaResponse, *entity.Error) {
	return u.socialRepo.GetSocial(ctx)
}

func validateSocial(social *entity.SocialMedia) error {
	if social.Name == "" {
		return errors.New("Name cannot be empty")
	}
	if social.SocialMediaUrl == "" {
		return errors.New("Social media url cannot be empty")
	}
	return nil
}
