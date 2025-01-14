package user

import (
	"ideadeck/usecase/user"
)

type CreatePresenter struct{}

func NewCreatePresenter() *CreatePresenter {
	return &CreatePresenter{}
}

func (p *CreatePresenter) Output() user.CreateUserOutput {
	return user.CreateUserOutput{}
}
