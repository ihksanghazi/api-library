package middleware

import (
	"context"
	"net/http"
	"os"

	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/utils"
	"gorm.io/gorm"
)

type Middleware interface {
	ValidToken(next http.Handler) http.Handler
	IsAdmin(next http.Handler) http.Handler
}

type MiddlewareImpl struct {
	ctx   context.Context
	db    *gorm.DB
	model domain.User
}

func NewMiddleware(ctx context.Context, db *gorm.DB, model domain.User) Middleware {
	return &MiddlewareImpl{
		ctx:   ctx,
		db:    db,
		model: model,
	}
}

func (m *MiddlewareImpl) ValidToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//ambil header Access-Token
		accessToken := r.Header.Get("Access-Token")
		if accessToken == "" {
			utils.ResponseError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		//parsing token
		_, errParsing := utils.ParsingToken(accessToken, os.Getenv("ACCESS_TOKEN_JWT"))
		if errParsing != nil {
			utils.ResponseError(w, http.StatusUnauthorized, errParsing.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *MiddlewareImpl) IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//get refresh token from cookie
		refreshToken, err := r.Cookie("AccessToken")
		if err != nil {
			utils.ResponseError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		// get user by refresh token
		errQuery := m.db.Model(m.model).WithContext(m.ctx).Where("refresh_token = ?", refreshToken.Value).First(&m.model).Error
		if errQuery != nil {
			utils.ResponseError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		// if user not admin return error
		if m.model.Role != "admin" {
			utils.ResponseError(w, http.StatusUnauthorized, "Not Admin")
			return
		}

		next.ServeHTTP(w, r)
	})
}
