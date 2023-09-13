package middleware

import (
	"net/http"
	"os"

	"github.com/ihksanghazi/api-library/repositories"
	"github.com/ihksanghazi/api-library/utils"
)

type Middleware interface {
	ValidToken(next http.Handler) http.Handler
	IsAdmin(next http.Handler) http.Handler
}

type MiddlewareImpl struct {
	repository *repositories.Query
}

func NewMiddleware(repository *repositories.Query) Middleware {
	return &MiddlewareImpl{
		repository: repository,
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
		user, errQuery := m.repository.User.Where(m.repository.User.RefreshToken.Eq(refreshToken.Value)).First()
		if errQuery != nil {
			utils.ResponseError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		// if user not admin return error
		if user.Role != "admin" {
			utils.ResponseError(w, http.StatusUnauthorized, "Not Admin")
			return
		}

		next.ServeHTTP(w, r)
	})
}
