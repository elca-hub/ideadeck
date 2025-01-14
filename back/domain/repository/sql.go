package repository

import "ideadeck/domain/repository/sql"

// 複数のinterfaceを1つにまとめる

type SQL interface {
	UserRepository() sql.UserRepository
}
