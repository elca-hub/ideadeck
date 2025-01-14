package user

import (
	"ideadeck/usecase/user"
)

type VerificationEmailPresenter struct{}

func NewVerificationEmailPresenter() *VerificationEmailPresenter {
	return &VerificationEmailPresenter{}
}

func (p *VerificationEmailPresenter) Output() user.VerificationEmailOutput {
	return user.VerificationEmailOutput{}
}
