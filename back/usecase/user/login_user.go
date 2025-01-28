package user

import (
	"github.com/go-playground/validator/v10"
	"ideadeck/domain/model"
	"ideadeck/domain/repository/nosql"
	"ideadeck/domain/repository/sql"
	"ideadeck/infra/auth"
)

type (
	LoginUserUseCase interface {
		Execute(LoginUserInput) (LoginUserOutput, error)
	}

	LoginUserInput struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=8,max=32"`
	}

	LoginUserPresenter interface {
		Output(user model.User, token string) LoginUserOutput
	}

	LoginUserOutput struct {
		Email string `json:"email"`
		Token string `json:"-"`
	}

	loginUserInterator struct {
		sqlRepository   sql.UserRepository
		noSqlRepository nosql.UserRepository
		presenter       LoginUserPresenter
	}
)

func NewLoginUserInterator(
	sqlRepository sql.UserRepository,
	noSqlRepository nosql.UserRepository,
	presenter LoginUserPresenter,
) LoginUserUseCase {
	return loginUserInterator{
		sqlRepository:   sqlRepository,
		noSqlRepository: noSqlRepository,
		presenter:       presenter,
	}
}

func (i loginUserInterator) Execute(input LoginUserInput) (LoginUserOutput, error) {
	validate := validator.New()

	if err := validate.Struct(input); err != nil {
		return i.presenter.Output(model.User{}, ""), err
	}

	inputEmail, err := model.NewEmail(input.Email)

	if err != nil {
		return i.presenter.Output(model.User{}, ""), err
	}

	userModel, err := i.sqlRepository.FindByEmail(inputEmail)

	if err != nil {
		return i.presenter.Output(model.User{}, ""), err
	}

	if !auth.CheckPasswordHash(input.Password, userModel.Password()) {
		return i.presenter.Output(*userModel, ""), nil
	}

	session, err := i.noSqlRepository.StartSession(inputEmail)

	if err != nil {
		return i.presenter.Output(*userModel, ""), err
	}

	return i.presenter.Output(*userModel, session), nil
}
