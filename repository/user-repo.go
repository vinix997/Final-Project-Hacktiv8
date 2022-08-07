package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/vinix997/Final-Project-Hacktiv8/entity"
	"github.com/vinix997/Final-Project-Hacktiv8/helper"
	"github.com/vinix997/Final-Project-Hacktiv8/service"
)

type userRepo struct {
	sql *sql.DB
}

func NewUserRepository(sql *sql.DB) service.UserRepository {
	return &userRepo{sql: sql}
}

func (o *userRepo) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	bytes, err := helper.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	query := "INSERT INTO USER (USERNAME, EMAIL, PASSWORD, AGE, CREATED_AT) VALUES(?,?,?,?,?)"
	row, errQuery := o.sql.ExecContext(ctx, query, user.Username, user.Email, bytes, user.Age, time.Now())
	if errQuery != nil {
		return nil, errQuery
	}
	id, err := row.LastInsertId()
	user.Id = int(id)
	return user, nil
}

func (o *userRepo) UpdateUser(ctx context.Context, user *entity.User, id int) (*entity.User, *entity.Error) {
	query := "UPDATE USER SET EMAIL = ?, USERNAME = ?, UPDATED_AT = ?  WHERE ID = ? "
	_, err := o.sql.ExecContext(ctx, query, user.Email, user.Username, time.Now(), id)
	if err != nil {
		var err *entity.Error
		err.IsError = true
		err.Message = "Failed to update"
		return nil, err
	}
	return user, nil
}

func (o *userRepo) Login(ctx context.Context, user *entity.User) (string, *entity.Error) {
	var tempUser entity.User
	query := "SELECT ID, PASSWORD FROM USER WHERE EMAIL = ?"
	row := o.sql.QueryRowContext(ctx, query, user.Email)
	row.Scan(&tempUser.Id, &tempUser.Password)

	check := helper.CheckPasswordHash(user.Password, tempUser.Password)
	if !check {
		var err *entity.Error
		helper.SetError(err, "Incorrect email/password !")
		return "", err
	}
	validToken, err := helper.GenerateJWT(tempUser.Id)
	if err != nil {
		var err *entity.Error
		err = helper.SetError(err, "Token failure !")
		return "", err
	}
	return validToken, nil
}

func (o *userRepo) DeleteUser(ctx context.Context, id int) (string, *entity.Error) {
	query := "DELETE FROM USER WHERE ID = ? "
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
	message := "Your account has been successfully deleted"
	return message, nil
}
