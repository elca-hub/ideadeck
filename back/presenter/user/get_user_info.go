package user

import (
	usermodel "ideadeck/domain/model"
	"ideadeck/usecase/user"
)

type GetUserInfoPresenter struct{}

func NewGetUserInfoPresenter() *GetUserInfoPresenter {
	return &GetUserInfoPresenter{}
}

func (p *GetUserInfoPresenter) Output(model usermodel.User, token string) user.GetUserInfoOutput {
	return user.GetUserInfoOutput{
		Email: model.Email().Email(),
		Token: token,
	}
}
