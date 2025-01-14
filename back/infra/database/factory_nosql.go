package database

import (
	"errors"
	"ideadeck/domain/repository"
)

const (
	InstanceRedis int = iota
)

func NewDatabaseNoSqlFactory(instance int) (repository.NoSQL, error) {
	switch instance {
	case InstanceRedis:
		return NewRedisHandler(NewMyNoSQLConfig())
	default:
		return nil, errors.New("invalid instance")
	}
}
