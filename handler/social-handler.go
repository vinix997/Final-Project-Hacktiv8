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

type SocialHandlerInterface interface {
	SocialHandler(w http.ResponseWriter, r *http.Request)
}

type SocialHandler struct {
	socialService service.SocialService
	sql           *sql.DB
}

func NewSocialHandler(socialService service.SocialService, sql *sql.DB) SocialHandlerInterface {
	return &SocialHandler{socialService: socialService, sql: sql}
}

func (h *SocialHandler) SocialHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	paramId := params["socialMediaId"]
	if r.Method == "POST" {
		h.createSocialHandler(w, r)
	}
	if r.Method == "PUT" {
		id, err := strconv.Atoi(paramId)
		if err != nil {
			helper.WriteJsonResponse(w, 504, err)
		}
		h.updateSocialHandler(w, r, id)
	}
	if r.Method == "DELETE" {
		id, err := strconv.Atoi(paramId)
		if err != nil {
			helper.WriteJsonResponse(w, 504, err)
		}
		h.deleteSocialHandler(w, r, id)
	}
	if r.Method == "GET" {
		h.getSocialHandler(w, r)
	}
}

func (h *SocialHandler) createSocialHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var social entity.SocialMedia
	if err := decoder.Decode(&social); err != nil {
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
	data, repErr := h.socialService.CreateSocial(r.Context(), &entity.SocialMedia{
		Name:           social.Name,
		SocialMediaUrl: social.SocialMediaUrl,
		UserId:         userId,
	})
	if repErr != nil {
		var errorResponse entity.Error
		helper.SetError(&errorResponse, "Error creating user")
		helper.WriteJsonResponse(w, 504, errorResponse)
		return
	}
	helper.WriteJsonResponse(w, 201, data)
}

func (h *SocialHandler) updateSocialHandler(w http.ResponseWriter, r *http.Request, id int) {
	decoder := json.NewDecoder(r.Body)
	var social entity.SocialMedia
	if err := decoder.Decode(&social); err != nil {
		helper.WriteJsonResponse(w, 504, err)
		return
	}

	updated, upErr := h.socialService.UpdateSocial(r.Context(), &entity.SocialMedia{
		SocialMediaUrl: social.SocialMediaUrl,
		Name:           social.Name,
	}, id)
	if upErr != nil {
		helper.WriteJsonResponse(w, 504, upErr)
		return
	}
	updated.Id = id
	helper.WriteJsonResponse(w, 200, updated)
}

func (h *SocialHandler) getSocialHandler(w http.ResponseWriter, r *http.Request) {
	data, err := h.socialService.GetSocial(r.Context())
	if err != nil {
		helper.WriteJsonResponse(w, 504, err)
	}
	helper.WriteJsonResponse(w, 200, data)
}

func (h *SocialHandler) deleteSocialHandler(w http.ResponseWriter, r *http.Request, id int) {
	delete, err := h.socialService.DeleteSocial(r.Context(), id)
	if err != nil {
		helper.WriteJsonResponse(w, 504, err)
		return
	}

	helper.WriteJsonResponse(w, 200, delete)
}
