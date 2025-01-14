package user

import (
	"ideadeck/domain/model"
	"ideadeck/usecase/user"
)

type CreatePresenter struct{}

func NewCreatePresenter() *CreatePresenter {
	return &CreatePresenter{}
}

func (p *CreatePresenter) Output(account model.User) user.CreateUserOutput {
	return user.CreateUserOutput{}
}
