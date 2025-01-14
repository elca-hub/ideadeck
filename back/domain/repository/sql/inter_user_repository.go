package sql

import "ideadeck/domain/model"

type UserRepository interface {
	Create(u *model.User) error
	Exists(email *model.Email) (bool, error)
}
