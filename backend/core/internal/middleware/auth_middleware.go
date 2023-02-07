package middleware

import (
	"backend/core/helper"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unable to authenticate"))
			return
		}

		uc, err := helper.AnalyzeToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unable to authenticate"))
			return
		}
		r.Header.Set("user_id", string(rune(uc.Id)))
		r.Header.Set("user_identity", uc.Identity)
		r.Header.Set("user_name", uc.Name)
		next(w, r)
	}
}
