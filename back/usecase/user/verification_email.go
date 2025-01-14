package user

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"ideadeck/domain/model"
	"ideadeck/domain/repository/nosql"
	"ideadeck/domain/repository/sql"
)

type (
	VerificationEmailUseCase interface {
		Execute(VerificationEmailInput) (VerificationEmailOutput, error)
	}

	VerificationEmailInput struct {
		Token string `validate:"required"`
	}

	VerificationEmailPresenter interface {
		Output() VerificationEmailOutput
	}

	VerificationEmailOutput struct {
	}

	verificationEmailInterator struct {
		sqlRepository   sql.UserRepository
		noSqlRepository nosql.UserRepository
		presenter       VerificationEmailPresenter
	}
)

func NewVerificationEmailInterator(
	sqlRepository sql.UserRepository,
	noSqlRepository nosql.UserRepository,
	presenter VerificationEmailPresenter,
) VerificationEmailUseCase {
	return verificationEmailInterator{
		sqlRepository:   sqlRepository,
		noSqlRepository: noSqlRepository,
		presenter:       presenter,
	}
}

func (i verificationEmailInterator) Execute(input VerificationEmailInput) (VerificationEmailOutput, error) {
	validate := validator.New()

	if err := validate.Struct(input); err != nil {
		return VerificationEmailOutput{}, err
	}

	userEmail, err := i.noSqlRepository.GetSession(input.Token)

	if err != nil {
		return VerificationEmailOutput{}, errors.New("invalid token")
	}

	userModel, err := i.sqlRepository.FindByEmail(userEmail)

	if err != nil {
		return VerificationEmailOutput{}, err
	}

	if userModel.EmailVerification() != model.InConfirmation {
		return VerificationEmailOutput{}, errors.New("already confirmed")
	}

	userModel.UpdateEmailVerification(model.Confirmed)

	if err := i.sqlRepository.Update(userModel); err != nil {
		return VerificationEmailOutput{}, err
	}

	if _, err := i.noSqlRepository.StartSession(userEmail); err != nil {
		return VerificationEmailOutput{}, err
	}

	return VerificationEmailOutput{}, nil
}
