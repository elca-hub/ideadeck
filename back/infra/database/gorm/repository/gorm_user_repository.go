package repository

import (
	"gorm.io/gorm"
	"ideadeck/domain/model"
	"ideadeck/infra/database/gorm/gorm_model"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (r GormUserRepository) Create(user *model.User) error {
	email := user.Email()

	gormUser := gorm_model.User{
		ID:                user.ID().ID(),
		Email:             email.Email(),
		Password:          user.Password(),
		EmailVerification: user.EmailVerification(),
	}
	
	return r.db.Create(&gormUser).Error
}

func (r GormUserRepository) Exists(email *model.Email) (bool, error) {
	var counter int64

	r.db.Model(&gorm_model.User{}).Where("email = ?", email.Email()).Count(&counter)

	return counter > 0, nil
}
