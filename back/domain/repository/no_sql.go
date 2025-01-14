package repository

import "ideadeck/domain/repository/nosql"

type NoSQL interface {
	UserRepository() nosql.UserRepository
}
