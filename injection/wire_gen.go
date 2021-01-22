// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injection

import (
	"github.com/jinzhu/gorm"
	"github.com/teten-nugraha/golang-twitter-clone-api/api"
	"github.com/teten-nugraha/golang-twitter-clone-api/repository"
	"github.com/teten-nugraha/golang-twitter-clone-api/service"
)

// Injectors from wire.go:

func InitUserApi(db *gorm.DB) api.UserApi {
	userRepository := repository.ProviderUserRepository(db)
	userService := service.ProviderUserService(userRepository)
	userApi := api.ProviderUserAPI(userService)
	return userApi
}
