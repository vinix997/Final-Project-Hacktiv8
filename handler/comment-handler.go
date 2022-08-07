package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/vinix997/Final-Project-Hacktiv8/entity"
	"github.com/vinix997/Final-Project-Hacktiv8/helper"
	"github.com/vinix997/Final-Project-Hacktiv8/service"
)

type CommentHandlerInterface interface {
	CommentHandler(w http.ResponseWriter, r *http.Request)
}

type CommentHandler struct {
	commentService service.CommentService
	sql            *sql.DB
}

func NewCommentHandler(commentService service.CommentService, sql *sql.DB) CommentHandlerInterface {
	return &CommentHandler{commentService: commentService, sql: sql}
}

func (h *CommentHandler) CommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	paramId := params["commentId"]
	if r.Method == "POST" {
		h.createCommentHandler(w, r)
	}
	if r.Method == "PUT" {
		id, err := strconv.Atoi(paramId)
		if err != nil {
			helper.WriteJsonResponse(w, 504, err)
		}
		h.updateCommentHandler(w, r, id)
	}
	if r.Method == "DELETE" {
		id, err := strconv.Atoi(paramId)
		if err != nil {
			helper.WriteJsonResponse(w, 504, err)
		}
		h.deleteCommentHandler(w, r, id)
	}
	if r.Method == "GET" {
		h.getCommentHandler(w, r)
	}
}

func (h *CommentHandler) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var comment entity.Comment
	if err := decoder.Decode(&comment); err != nil {
		helper.WriteJsonResponse(w, 504, err)
		return
	}

	tokenString, ok := helper.ExtractTokenFromAuthHeader(r.Header.Get("Authorization"))
	if !ok {
		var err entity.Error
		err.Message = "Invalid token"
		helper.WriteJsonResponse(w, 504, err)
		return
	}
	claims := jwt.MapClaims{}
	_, lol := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(helper.GetEnvVar("SECRET_KEY")), nil
	})
	if lol != nil {
		var err entity.Error
		err.Message = "Unexpected error when parsing token"
		helper.WriteJsonResponse(w, 504, err)
		return
	}
	userId := int(claims["userId"].(float64))
	data, repErr := h.commentService.CreateComment(r.Context(), &entity.Comment{
		Message: comment.Message,
		PhotoId: comment.PhotoId,
		UserId:  userId,
	})
	if repErr != nil {
		var errorResponse entity.Error
		helper.SetError(&errorResponse, "Error creating comment")
		helper.WriteJsonResponse(w, 504, errorResponse)
		return
	}
	helper.WriteJsonResponse(w, 201, data)
}

func (h *CommentHandler) updateCommentHandler(w http.ResponseWriter, r *http.Request, id int) {
	decoder := json.NewDecoder(r.Body)
	var comment entity.Comment
	if err := decoder.Decode(&comment); err != nil {
		helper.WriteJsonResponse(w, 504, err)
		return
	}

	updated, upErr := h.commentService.UpdateComment(r.Context(), &entity.Comment{
		Message: comment.Message,
	}, id)
	if upErr != nil {
		helper.WriteJsonResponse(w, 504, upErr)
		return
	}
	updated.Id = id
	helper.WriteJsonResponse(w, 200, updated)
}

func (h *CommentHandler) getCommentHandler(w http.ResponseWriter, r *http.Request) {
	data, err := h.commentService.GetComment(r.Context())
	if err != nil {
		helper.WriteJsonResponse(w, 504, err)
	}
	helper.WriteJsonResponse(w, 200, data)
}

func (h *CommentHandler) deleteCommentHandler(w http.ResponseWriter, r *http.Request, id int) {
	delete, err := h.commentService.DeleteComment(r.Context(), id)
	if err != nil {
		helper.WriteJsonResponse(w, 504, err)
		return
	}

	helper.WriteJsonResponse(w, 200, delete)
}
