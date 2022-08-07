package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vinix997/Final-Project-Hacktiv8/entity"
	"github.com/vinix997/Final-Project-Hacktiv8/helper"
	"github.com/vinix997/Final-Project-Hacktiv8/service"
)

type UserHandlerInterface interface {
	UserHandler(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	userService service.UserService
	sql         *sql.DB
}

func NewUserHandler(userService service.UserService, sql *sql.DB) UserHandlerInterface {
	return &UserHandler{userService: userService, sql: sql}
}
func (h *UserHandler) UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	paramId := params["userId"]

	if r.Method == "POST" {
		if r.URL.Path == "/users/register" {
			h.createUserHandler(w, r)
		} else if r.URL.Path == "/users/login" {
			h.loginUserHandler(w, r)
		}
	}
	if r.Method == "PUT" {
		id, err := strconv.Atoi(paramId)
		if err != nil {
			helper.WriteJsonResponse(w, 504, err)
		}
		h.updateUserHandler(w, r, id)
	}
	if r.Method == "DELETE" {
		id, err := strconv.Atoi(paramId)
		if err != nil {
			helper.WriteJsonResponse(w, 504, err)
		}
		h.deleteUserHandler(w, r, id)
	}

}

func (h *UserHandler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user entity.User
	if err := decoder.Decode(&user); err != nil {
		helper.WriteJsonResponse(w, 504, err)
		return
	}

	data, repErr := h.userService.CreateUser(r.Context(), &entity.User{
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
		Id:       user.Id,
		Password: user.Password,
	})
	if repErr != nil {
		var errorResponse entity.Error
		helper.SetError(&errorResponse, "Error creating user")
		helper.WriteJsonResponse(w, 504, errorResponse)
		return
	}

	helper.WriteJsonResponse(w, 201, data)
}

func (h *UserHandler) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	var authDetails entity.Authentication
	err := json.NewDecoder(r.Body).Decode(&authDetails)

	if err != nil {
		var err entity.Error
		helper.SetError(&err, "Error in reading auth")
		helper.WriteJsonResponse(w, 504, err)
		return
	}
	tokenString, loginErr := h.userService.Login(r.Context(), &entity.User{
		Email:    authDetails.Email,
		Password: authDetails.Password,
	})

	if loginErr != nil {
		helper.WriteJsonResponse(w, 504, loginErr)
		return
	}

	var token entity.Token
	token.TokenString = tokenString

	helper.WriteJsonResponse(w, 200, token)
}

func (h *UserHandler) updateUserHandler(w http.ResponseWriter, r *http.Request, id int) {
	decoder := json.NewDecoder(r.Body)
	var user entity.User
	if err := decoder.Decode(&user); err != nil {
		helper.WriteJsonResponse(w, 504, err)
		return
	}

	var tempUser entity.User
	query := "SELECT id from user where id = ?"
	errQuery := h.sql.QueryRowContext(r.Context(), query, id).Scan(&tempUser.Id)
	if errQuery == sql.ErrNoRows {
		var err entity.Error
		helper.SetError(&err, "Unexpected error occured")
		helper.WriteJsonResponse(w, 504, err)
		return
	}
	updated, err := h.userService.UpdateUser(r.Context(), &entity.User{
		Username: user.Username,
		Email:    user.Email,
	}, id)
	if err != nil {
		helper.WriteJsonResponse(w, 504, err)
		return
	}

	responseUser := entity.User{
		Id:       id,
		Username: updated.Username,
		Age:      updated.Age,
		Email:    updated.Email,
	}

	helper.WriteJsonResponse(w, 200, responseUser)
}

func (h *UserHandler) deleteUserHandler(w http.ResponseWriter, r *http.Request, id int) {
	delete, err := h.userService.DeleteUser(r.Context(), id)
	if err != nil {
		helper.WriteJsonResponse(w, 504, err)
		return
	}

	helper.WriteJsonResponse(w, 200, delete)
}
