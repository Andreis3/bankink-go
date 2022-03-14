package app

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/santos/banking-go/domain"
	"github.com/santos/banking-go/errs"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (a AuthMiddleware) authorizationHandler() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			customerRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {
				token := getTokenFromHeader(authHeader)

				isAuthorized := a.repo.IsAuthorized(token, customerRoute.GetName(), currentRouteVars)

				if isAuthorized {
					next.ServeHTTP(w, r)
				} else {
					appError := errs.AppError{http.StatusForbidden, "Unauthorized"}
					writeResponse(w, appError.Code, appError.AsMessage())
				}
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	/*
	   token is coming in the format as below
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/

	spliToken := strings.Split(header, "Bearer")
	if len(spliToken) == 2 {
		return strings.TrimSpace(spliToken[1])
	}
	return ""
}
