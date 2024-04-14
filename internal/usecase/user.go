package usecase

import "github.com/mvp-mogila/avito-intership-backend-2024/internal/config"

type UserUsecase struct {
	adminToken string
	userToken  string
}

func NewUserUsecase(cfg config.AuthConfig) *UserUsecase {
	return &UserUsecase{
		adminToken: cfg.AdminToken,
		userToken:  cfg.UserToken,
	}
}

func (u *UserUsecase) CheckAdmin(token string) bool {
	return u.adminToken == token
}

func (u *UserUsecase) CheckUser(token string) bool {
	return u.userToken == token
}
