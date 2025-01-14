package nosql

import "ideadeck/domain/model"

type UserRepository interface {
	StartSession(email *model.Email) error
	GetSession(email *model.Email) (string, error)
	DeleteSession(email *model.Email) error
}
