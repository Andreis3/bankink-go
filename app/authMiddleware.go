package app

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/santos/banking-go/domain"
	"github.com/santos/banking-go/errs"
	"github.com/santos/banking-go/service"
)

type AuthMiddleware struct {
	service    AuthHandler
	permission domain.RolePermissions
}

func (a AuthMiddleware) authorizationHandler() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			customerRoute := mux.CurrentRoute(r)
			//currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {
				token := getTokenFromHeader(authHeader)

				parserToken := map[string]string{
					"token":      token,
					"route_name": customerRoute.GetName(),
				}

				isAuthorized := a.service.service.Verify(parserToken)

				if isAuthorized != nil {
					appError := errs.AppError{http.StatusForbidden, "Unauthorized"}
					writeResponse(w, appError.Code, appError.AsMessage())

				} else {
					next.ServeHTTP(w, r)
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

type DefaultAuthService struct {
	service    service.AuthService
	permission domain.RolePermissions
}

func NewAuthService() *DefaultAuthService {
	return &DefaultAuthService{
		permission: domain.GetRolePermissions(),
	}
}
