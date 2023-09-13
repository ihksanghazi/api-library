package middleware

import (
	"context"
	"net/http"
	"os"

	"github.com/ihksanghazi/api-library/utils"
)

type Middleware interface {
	ValidToken(next http.Handler) http.Handler
}

type MiddlewareImpl struct{}

func NewMiddleware() Middleware {
	return &MiddlewareImpl{}
}

type contextKey string

func (m *MiddlewareImpl) ValidToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//ambil header Access-Token
		accessToken := r.Header.Get("Access-Token")
		if accessToken == "" {
			utils.ResponseError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		//parsing token
		claims, errParsing := utils.ParsingToken(accessToken, os.Getenv("ACCESS_TOKEN_JWT"))
		if errParsing != nil {
			utils.ResponseError(w, http.StatusUnauthorized, errParsing.Error())
			return
		}

		ctx := context.WithValue(r.Context(), contextKey("id"), claims.ID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
