package user

import (
	"ideadeck/usecase/user"
)

type VerificationEmailPresenter struct{}

func NewVerificationEmailPresenter() *VerificationEmailPresenter {
	return &VerificationEmailPresenter{}
}

func (p *VerificationEmailPresenter) Output(token string) user.VerificationEmailOutput {
	return user.VerificationEmailOutput{
		Token: token,
	}
}
