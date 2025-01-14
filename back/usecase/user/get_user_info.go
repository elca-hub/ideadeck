package user

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"ideadeck/domain/model"
	"ideadeck/domain/repository/nosql"
	"ideadeck/domain/repository/sql"
)

type (
	GetUserInfoUseCase interface {
		Execute(GetUserInfoInput) (GetUserInfoOutput, error)
	}

	GetUserInfoInput struct {
		Token string `validate:"required"`
	}

	GetUserInfoPresenter interface {
		Output(user model.User) GetUserInfoOutput
	}

	GetUserInfoOutput struct {
		Email string `json:"email"`
	}

	getUserInfoInterator struct {
		sqlRepository   sql.UserRepository
		noSqlRepository nosql.UserRepository
		presenter       GetUserInfoPresenter
	}
)

func NewGetUserInfoInterator(
	sqlRepository sql.UserRepository,
	noSqlRepository nosql.UserRepository,
	presenter GetUserInfoPresenter,
) GetUserInfoUseCase {
	return getUserInfoInterator{
		sqlRepository:   sqlRepository,
		noSqlRepository: noSqlRepository,
		presenter:       presenter,
	}
}

func (i getUserInfoInterator) Execute(input GetUserInfoInput) (GetUserInfoOutput, error) {
	validate := validator.New()

	if err := validate.Struct(input); err != nil {
		return i.presenter.Output(model.User{}), err
	}

	userEmail, err := i.noSqlRepository.GetSession(input.Token)

	if err != nil {
		return i.presenter.Output(model.User{}), errors.New("invalid token")
	}

	userModel, err := i.sqlRepository.FindByEmail(userEmail)

	if err != nil {
		return i.presenter.Output(model.User{}), err
	}

	return i.presenter.Output(*userModel), nil
}
