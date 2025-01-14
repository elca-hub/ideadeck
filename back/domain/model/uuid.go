package model

import (
	gouuid "github.com/satori/go.uuid"
)

type UUID struct {
	id string
}

func NewUUID() UUID {
	return UUID{
		id: gouuid.NewV4().String(),
	}
}

func (u UUID) ID() string {
	return u.id
}

func IsValidUUID(uuid string) bool {
	_, err := gouuid.FromString(uuid)

	return err == nil
}
