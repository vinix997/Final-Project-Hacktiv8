package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/vinix997/Final-Project-Hacktiv8/entity"
	"github.com/vinix997/Final-Project-Hacktiv8/service"
)

type photoRepo struct {
	sql *sql.DB
}

func NewPhotoRepository(sql *sql.DB) service.PhotoRepository {
	return &photoRepo{sql: sql}
}

func (o *photoRepo) CreatePhoto(ctx context.Context, photo *entity.Photo) (*entity.Photo, *entity.Error) {
	query := "INSERT INTO PHOTO (TITLE, CAPTION, PHOTO_URL, USER_ID, CREATED_AT, UPDATED_AT) VALUES(?,?,?,?,?,?)"
	res, errQuery := o.sql.ExecContext(ctx, query, photo.Title, photo.Caption, photo.PhotoUrl, photo.UserId, time.Now(), time.Now())
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
	photo.Id = int(id)
	return photo, nil
}

func (o *photoRepo) UpdatePhoto(ctx context.Context, photo *entity.Photo, id int) (*entity.Photo, *entity.Error) {
	query := "UPDATE PHOTO SET TITLE = ?, CAPTION = ?, PHOTO_URL = ?, UPDATED_AT = ?  WHERE id = ? "
	_, err := o.sql.ExecContext(ctx, query, photo.Title, photo.Caption, photo.PhotoUrl, time.Now(), id)
	if err != nil {
		var err *entity.Error
		err.IsError = true
		err.Message = "Failed to update"
		return nil, err
	}

	return photo, nil
}

func (o *photoRepo) GetPhoto(ctx context.Context) ([]entity.PhotoResponse, *entity.Error) {
	var photoResponses []entity.PhotoResponse
	query := "SELECT P.ID, P.TITLE, P.CAPTION, P.PHOTO_URL, P.USER_ID, P.CREATED_AT, P.UPDATED_AT, U.USERNAME, U.PASSWORD FROM PHOTO P JOIN USER U ON P.USER_ID = U.ID"
	rows, err := o.sql.QueryContext(ctx, query)
	if err != nil {
		var err *entity.Error
		err.IsError = true
		err.Message = "Error retrieving data"
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var photoResponse entity.PhotoResponse
		err := rows.Scan(&photoResponse.Id, &photoResponse.Title,
			&photoResponse.Caption, &photoResponse.PhotoUrl,
			&photoResponse.UserId, &photoResponse.CreatedAt,
			&photoResponse.UpdatedAt, &photoResponse.User.Username,
			&photoResponse.User.Password,
		)
		if err != nil {
			var err *entity.Error
			err.IsError = true
			err.Message = "Error retrieving data"
			return nil, err
		}
		photoResponses = append(photoResponses, photoResponse)
	}
	return photoResponses, nil
}

func (o *photoRepo) DeletePhoto(ctx context.Context, id int) (string, *entity.Error) {
	query := "DELETE FROM PHOTO WHERE ID = ?"
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
	message := "Your photo has been successfully deleted"
	return message, nil
}
