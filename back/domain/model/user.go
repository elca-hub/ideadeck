package model

import "time"

const (
	Unconfirmed int = iota
	InConfirmation
	Confirmed
)

type User struct {
	id                UUID
	email             *Email
	password          string
	createdAt         time.Time
	updatedAt         time.Time
	emailVerification int
	confirmedAt       time.Time
}

func NewUser(
	id UUID,
	email *Email,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
	emailVerification int,
	confirmedAt time.Time,
) *User {
	return &User{
		id,
		email,
		password,
		createdAt,
		updatedAt,
		emailVerification,
		confirmedAt,
	}
}

func (u *User) ID() UUID {
	return u.id
}

func (u *User) Email() *Email {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) EmailVerification() int {
	return u.emailVerification
}

func (u *User) ConfirmedAt() time.Time {
	return u.confirmedAt
}
