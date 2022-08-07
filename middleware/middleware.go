package middleware

import (
	"net/http"

	"github.com/vinix997/Final-Project-Hacktiv8/entity"
	"github.com/vinix997/Final-Project-Hacktiv8/helper"

	"github.com/golang-jwt/jwt"
)

type AuthMiddlewareIface interface {
	AuthMiddleware(next http.Handler) http.Handler
}

type AuthMiddleware struct {
}

func NewAuthMiddleware() AuthMiddlewareIface {
	return &AuthMiddleware{}

}

func (m *AuthMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/users/register" || r.URL.Path == "/users/login" {
			next.ServeHTTP(w, r)
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
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(helper.GetEnvVar("SECRET_KEY")), nil
		})
		if err != nil {
			var err entity.Error
			err.Message = "Unexpected error when parsing token"
			helper.WriteJsonResponse(w, 504, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
