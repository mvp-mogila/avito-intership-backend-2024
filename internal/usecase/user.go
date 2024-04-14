package usecase

type UserUsecase struct {
	adminToken string
	userToken  string
}

func NewUserUsecase(admToken, usrToken string) *UserUsecase {
	return &UserUsecase{
		adminToken: admToken,
		userToken:  usrToken,
	}
}

func (u *UserUsecase) CheckAdmin(token string) bool {
	return u.adminToken == token
}

func (u *UserUsecase) CheckUser(token string) bool {
	return u.userToken == token
}
