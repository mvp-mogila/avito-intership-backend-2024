package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/models"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/pkg/utils"
)

type Users interface {
	CheckUser(token string) bool
	CheckAdmin(token string) bool
}

type AdminStatusKey struct{}

// TODO: get user banner - only for user
func Authentication(userUsecase Users) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("token")

			var ctx context.Context
			if userUsecase.CheckAdmin(token) {
				ctx = context.WithValue(context.Background(), AdminStatusKey{}, true)
			} else if userUsecase.CheckUser(token) {
				ctx = context.WithValue(context.Background(), AdminStatusKey{}, false)
			} else {
				utils.SendErrorResponse(w, http.StatusUnauthorized, "")
				return
			}

			rCtx := r.WithContext(ctx)
			handler.ServeHTTP(w, rCtx)
		})
	}
}

func Authorization(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		adminStatus, err := getAdminStatusFromCtx(r.Context())
		if err != nil {
			utils.SendErrorResponse(w, http.StatusUnauthorized, "")
			return
		}
		if !adminStatus {
			utils.SendErrorResponse(w, http.StatusForbidden, "")
			return
		}
		handler.ServeHTTP(w, r)
	}
}

func getAdminStatusFromCtx(ctx context.Context) (bool, error) {
	status, ok := ctx.Value(AdminStatusKey{}).(bool)
	if !ok {
		return false, models.ErrNoAuth
	}
	return status, nil
}
