package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/teten-nugraha/golang-twitter-clone-api/domain"
)

type UserRepositoryContract interface {

	FindByUsername(username string) (domain.User, error)
	SaveOrUpdate(user domain.User) (domain.User,error)
}

type UserRepository struct {
	DB *gorm.DB
}

func ProviderUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{
		DB: DB,
	}
}

func (u *UserRepository) FindByUsername(username string) (domain.User, error) {
	var user domain.User

	if err := u.DB.First(&user, "username=?", username).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepository) SaveOrUpdate(user domain.User) (domain.User, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}