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

type PhotoHandlerInterface interface {
	PhotoHandler(w http.ResponseWriter, r *http.Request)
}

type PhotoHandler struct {
	photoService service.PhotoService
	sql          *sql.DB
}

func NewPhotoHandler(photoService service.PhotoService, sql *sql.DB) PhotoHandlerInterface {
	return &PhotoHandler{photoService: photoService, sql: sql}
}

func (h *PhotoHandler) PhotoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	paramId := params["photoId"]
	if r.Method == "POST" {
		h.createPhotoHandler(w, r)
	}
	if r.Method == "PUT" {
		id, err := strconv.Atoi(paramId)
		if err != nil {
			helper.WriteJsonResponse(w, 504, err)
		}
		h.updatePhotoHandler(w, r, id)
	}
	if r.Method == "DELETE" {
		id, err := strconv.Atoi(paramId)
		if err != nil {
			helper.WriteJsonResponse(w, 504, err)
		}
		h.deletePhotoHandler(w, r, id)
	}
	if r.Method == "GET" {
		h.getPhotoHandler(w, r)
	}
}

func (h *PhotoHandler) createPhotoHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var photo entity.Photo
	if err := decoder.Decode(&photo); err != nil {
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
	a := int(claims["userId"].(float64))
	data, repErr := h.photoService.CreatePhoto(r.Context(), &entity.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserId:   a,
	})
	if repErr != nil {
		var errorResponse entity.Error
		helper.SetError(&errorResponse, "Error creating photo")
		helper.WriteJsonResponse(w, 504, errorResponse)
		return
	}
	helper.WriteJsonResponse(w, 201, data)
}

func (h *PhotoHandler) updatePhotoHandler(w http.ResponseWriter, r *http.Request, id int) {
	decoder := json.NewDecoder(r.Body)
	var photo entity.Photo
	if err := decoder.Decode(&photo); err != nil {
		helper.WriteJsonResponse(w, 504, err)
		return
	}

	updated, upErr := h.photoService.UpdatePhoto(r.Context(), &entity.Photo{
		Caption:  photo.Caption,
		Title:    photo.Title,
		PhotoUrl: photo.PhotoUrl,
	}, id)
	if upErr != nil {
		helper.WriteJsonResponse(w, 504, upErr)
		return
	}
	updated.Id = id
	helper.WriteJsonResponse(w, 200, updated)
}

func (h *PhotoHandler) getPhotoHandler(w http.ResponseWriter, r *http.Request) {
	data, err := h.photoService.GetPhoto(r.Context())
	if err != nil {
		helper.WriteJsonResponse(w, 504, err)
	}
	helper.WriteJsonResponse(w, 200, data)
}

func (h *PhotoHandler) deletePhotoHandler(w http.ResponseWriter, r *http.Request, id int) {
	delete, err := h.photoService.DeletePhoto(r.Context(), id)
	if err != nil {
		helper.WriteJsonResponse(w, 504, err)
		return
	}

	helper.WriteJsonResponse(w, 200, delete)
}
