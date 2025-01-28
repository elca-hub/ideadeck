package user

import (
	usermodel "ideadeck/domain/model"
	"ideadeck/usecase/user"
)

type LoginUserPresenter struct{}

func NewLoginUserPresenter() *LoginUserPresenter {
	return &LoginUserPresenter{}
}

func (p *LoginUserPresenter) Output(model usermodel.User, token string) user.LoginUserOutput {
	return user.LoginUserOutput{
		Email: model.Email().Email(),
		Token: token,
	}
}
