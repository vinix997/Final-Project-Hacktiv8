package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/vinix997/Final-Project-Hacktiv8/entity"
	"github.com/vinix997/Final-Project-Hacktiv8/service"
)

type commentRepo struct {
	sql *sql.DB
}

func NewCommentRepository(sql *sql.DB) service.CommentRepository {
	return &commentRepo{sql: sql}
}

func (o *commentRepo) CreateComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, *entity.Error) {
	query := "INSERT INTO COMMENT (MESSAGE, PHOTO_ID, USER_ID, CREATED_AT, UPDATED_AT) VALUES(?,?,?,?,?)"
	res, errQuery := o.sql.ExecContext(ctx, query, comment.Message, comment.PhotoId, comment.UserId, time.Now(), time.Now())
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
	comment.Id = int(id)
	return comment, nil
}

func (o *commentRepo) UpdateComment(ctx context.Context, comment *entity.Comment, id int) (*entity.CommentUpdateResponse, *entity.Error) {
	query := "UPDATE COMMENT SET MESSAGE = ?, UPDATED_AT = ?  WHERE id = ? "
	_, err := o.sql.ExecContext(ctx, query, comment.Message, time.Now(), id)
	if err != nil {
		var err *entity.Error
		err.IsError = true
		err.Message = "Failed to update"
		return nil, err
	}
	var customResponse *entity.CommentUpdateResponse
	getQuery := "SELECT C.ID, P.TITLE, P.CAPTION, P.PHOTO_URL, C.USER_ID, C.UPDATED_AT FROM COMMENT C JOIN PHOTO P ON C.USER_ID = P.ID WHERE C.ID = ? "
	_ = o.sql.QueryRowContext(ctx, getQuery, id).Scan(
		&customResponse.Id, &customResponse.Title,
		&customResponse.Caption, &customResponse.PhotoUrl,
		&customResponse.UserId, &customResponse.UpdatedAt,
	)

	return customResponse, nil
}

func (o *commentRepo) GetComment(ctx context.Context) ([]entity.CommentResponse, *entity.Error) {
	var commentResponses []entity.CommentResponse
	query := "SELECT C.ID, C.MESSAGE, C.CREATED_AT, C.UPDATED_AT, P.ID, P.TITLE, P.CAPTION, P.PHOTO_URL, P.USER_ID, P.CREATED_AT, P.UPDATED_AT, U.ID, U.USERNAME, U.PASSWORD FROM COMMENT C JOIN PHOTO P ON C.PHOTO_ID = P.ID JOIN USER U ON P.USER_ID = U.ID"
	rows, err := o.sql.QueryContext(ctx, query)
	if err != nil {
		var err *entity.Error
		err.IsError = true
		err.Message = "Error retrieving data"
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var commentResponse entity.CommentResponse
		err := rows.Scan(
			&commentResponse.Id, &commentResponse.Message,
			&commentResponse.CreatedAt, &commentResponse.UpdatedAt,
			&commentResponse.Photo.Id, &commentResponse.Photo.Title,
			&commentResponse.Photo.Caption, &commentResponse.Photo.PhotoUrl,
			&commentResponse.Photo.UserId, &commentResponse.Photo.CreatedAt,
			&commentResponse.Photo.UpdatedAt, &commentResponse.User.Id,
			&commentResponse.User.Username, &commentResponse.User.Password,
		)
		if err != nil {
			var err *entity.Error
			err.IsError = true
			err.Message = "Error retrieving data"
			return nil, err
		}
		commentResponses = append(commentResponses, commentResponse)
	}
	return commentResponses, nil
}

func (o *commentRepo) DeleteComment(ctx context.Context, id int) (string, *entity.Error) {
	query := "DELETE FROM COMMENT WHERE ID = ?"
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
	message := "Your comment has been successfully deleted"
	return message, nil
}
