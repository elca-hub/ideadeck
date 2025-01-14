package user

import (
	user_model "ideadeck/domain/model"
	"ideadeck/usecase/user"
)

type GetUserInfoPresenter struct{}

func NewGetUserInfoPresenter() *GetUserInfoPresenter {
	return &GetUserInfoPresenter{}
}

func (p *GetUserInfoPresenter) Output(model user_model.User) user.GetUserInfoOutput {
	return user.GetUserInfoOutput{
		Email: model.Email().Email(),
	}
}
