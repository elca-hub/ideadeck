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

func (r GormUserRepository) Update(user *model.User) error {
	email := user.Email()

	gormUser := gorm_model.User{
		ID:                user.ID().ID(),
		Email:             email.Email(),
		Password:          user.Password(),
		EmailVerification: user.EmailVerification(),
		CreatedAt:         user.CreatedAt(),
	}

	return r.db.Save(&gormUser).Error
}

func (r GormUserRepository) FindByEmail(email *model.Email) (*model.User, error) {
	var gormUser gorm_model.User

	if err := r.db.Where("email = ?", email.Email()).First(&gormUser).Error; err != nil {
		return nil, err
	}

	userEmail, err := model.NewEmail(gormUser.Email)

	if err != nil {
		return nil, err
	}

	user := model.NewUser(
		model.NewUUID(gormUser.ID),
		userEmail,
		gormUser.Password,
		gormUser.CreatedAt,
		gormUser.UpdatedAt,
		gormUser.EmailVerification,
	)

	return user, nil
}

func (r GormUserRepository) FetchInConfirmationUsers() ([]*model.User, error) {
	var gormUsers []gorm_model.User

	if err := r.db.Where("email_verification = ?", model.InConfirmation).Find(&gormUsers).Error; err != nil {
		return nil, err
	}

	var users []*model.User

	for _, gormUser := range gormUsers {
		userEmail, err := model.NewEmail(gormUser.Email)

		if err != nil {
			return nil, err
		}

		user := model.NewUser(
			model.NewUUID(gormUser.ID),
			userEmail,
			gormUser.Password,
			gormUser.CreatedAt,
			gormUser.UpdatedAt,
			gormUser.EmailVerification,
		)

		users = append(users, user)
	}

	return users, nil
}
