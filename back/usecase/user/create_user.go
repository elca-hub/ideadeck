package user

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"ideadeck/domain/model"
	"ideadeck/domain/repository/nosql"
	"ideadeck/domain/repository/sql"
	"ideadeck/infra/auth"
	"ideadeck/infra/email"
	"time"
)

type (
	CreateUserUserCase interface {
		Execute(CreateUserInput) (CreateUserOutput, error)
	}

	CreateUserInput struct {
		Name     string `validate:"required"`
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=8,max=32"`
	}

	CreateUserPresenter interface {
		Output() CreateUserOutput
	}

	CreateUserOutput struct {
	}

	createUserInterator struct {
		sqlRepository   sql.UserRepository
		noSqlRepository nosql.UserRepository
		presenter       CreateUserPresenter
	}
)

func NewCreateUserInterator(
	sqlRepository sql.UserRepository,
	noSqlRepository nosql.UserRepository,
	presenter CreateUserPresenter,
) CreateUserUserCase {
	return createUserInterator{
		sqlRepository:   sqlRepository,
		noSqlRepository: noSqlRepository,
		presenter:       presenter,
	}
}

func (i createUserInterator) Execute(input CreateUserInput) (CreateUserOutput, error) {
	hashedPw := auth.HashPassword(input.Password)

	validate := validator.New()

	if err := validate.Struct(input); err != nil {
		return i.presenter.Output(), err
	}

	userEmail, err := model.NewEmail(input.Email)

	if err != nil {
		return i.presenter.Output(), err
	}

	isExists, err := i.sqlRepository.Exists(userEmail) // ユーザが存在するか確認

	if err != nil {
		return i.presenter.Output(), err
	}
	if isExists {
		return i.presenter.Output(), errors.New("user already exists")
	}

	user := model.NewUser(model.NewUUID(""), userEmail, hashedPw, time.Now(), time.Now(), model.InConfirmation)

	if err := i.sqlRepository.Create(user); err != nil {
		return i.presenter.Output(), err
	}

	token, err := i.noSqlRepository.StartSession(userEmail)

	if err != nil {
		return i.presenter.Output(), err
	}

	mailSubject := "【メール確認のお願い】"

	mailContent := fmt.Sprintf("以下のリンクをクリックしてメールアドレスを確認してください。\nhttp://localhost:8080/verification/email?token=%s", token)

	if err := email.SmtpSendMail([]string{input.Email}, mailSubject, mailContent); err != nil {
		return i.presenter.Output(), err
	}

	return i.presenter.Output(), nil
}
