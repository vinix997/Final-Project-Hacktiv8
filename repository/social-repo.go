package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/vinix997/Final-Project-Hacktiv8/entity"
	"github.com/vinix997/Final-Project-Hacktiv8/service"
)

type socialRepo struct {
	sql *sql.DB
}

func NewSocialRepository(sql *sql.DB) service.SocialRepository {
	return &socialRepo{sql: sql}
}

func (o *socialRepo) CreateSocial(ctx context.Context, social *entity.SocialMedia) (*entity.SocialMedia, *entity.Error) {
	fmt.Println(social)
	query := "INSERT INTO SOCIALMEDIA (NAME, SOCIAL_MEDIA_URL, USERID, CREATED_AT, UPDATED_AT) VALUES(?,?,?,?,?)"
	res, errQuery := o.sql.ExecContext(ctx, query, social.Name, social.SocialMediaUrl, social.UserId, time.Now(), time.Now())
	if errQuery != nil {
		var errModel *entity.Error
		errModel.IsError = true
		errModel.Message = errQuery.Error()
		return nil, errModel
	}
	id, err := res.LastInsertId()
	if err != nil {
		var errModel *entity.Error
		errModel.IsError = true
		errModel.Message = err.Error()
		return nil, errModel
	}
	social.Id = int(id)
	return social, nil
}

func (o *socialRepo) UpdateSocial(ctx context.Context, social *entity.SocialMedia, id int) (*entity.SocialMedia, *entity.Error) {
	query := "UPDATE SOCIALMEDIA SET NAME = ?, SOCIAL_MEDIA_URL = ?, UPDATED_AT = ?  WHERE ID = ? "
	_, err := o.sql.ExecContext(ctx, query, social.Name, social.SocialMediaUrl, time.Now(), id)
	if err != nil {
		fmt.Println(social)
		var err *entity.Error
		err.IsError = true
		err.Message = "Failed to update"
		return nil, err
	}
	return social, nil
}

func (o *socialRepo) DeleteSocial(ctx context.Context, id int) (string, *entity.Error) {
	query := "DELETE FROM SOCIALMEDIA WHERE ID = ?"
	t, err := o.sql.ExecContext(ctx, query, id)
	_, rowError := t.LastInsertId()
	if rowError != nil {
		var err *entity.Error
		err.Message = rowError.Error()
		return "", err
	}
	if err != nil {
		var err *entity.Error
		err.Message = "Failed to delete"
		return "", err
	}
	message := "Your social media has been successfully deleted"
	return message, nil
}

func (o *socialRepo) GetSocial(ctx context.Context) (*entity.CustomSocialMediaResponse, *entity.Error) {
	var socmeds entity.CustomSocialMediaResponse
	query := "SELECT S.ID, S.NAME, S.SOCIAL_MEDIA_URL, S.USERID, S.CREATED_AT, S.UPDATED_AT, U.ID, U.USERNAME, U.PASSWORD FROM SOCIALMEDIA S JOIN USER U ON S.USERID = U.ID"
	rows, err := o.sql.QueryContext(ctx, query)
	if err != nil {
		var err *entity.Error
		err.IsError = true
		err.Message = "Error retrieving data"
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var socialmedia entity.SocialMediaResponse
		err := rows.Scan(&socialmedia.Id, &socialmedia.Name,
			&socialmedia.SocialMediaUrl, &socialmedia.UserId, &socialmedia.CreatedAt,
			&socialmedia.UpdatedAt, &socialmedia.User.Id,
			&socialmedia.User.Username, &socialmedia.User.Password)
		if err != nil {
			var err *entity.Error
			err.IsError = true
			err.Message = "Error retrieving data"
			return nil, err
		}
		socmeds.SocialMedias = append(socmeds.SocialMedias, socialmedia)
	}
	return &socmeds, nil
}
