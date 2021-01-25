package injection

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/teten-nugraha/golang-twitter-clone-api/api"
	"github.com/teten-nugraha/golang-twitter-clone-api/repository"
	"github.com/teten-nugraha/golang-twitter-clone-api/service"
)

func initUserApi(db *gorm.DB) api.UserApi {
	wire.Build(repository.ProviderUserRepository, service.ProviderUserService, api.ProviderUserAPI)

	return api.UserApi{}
}

func initAuthAPI(db *gorm.DB) api.AuthApi {
	wire.Build(repository.ProviderUserRepository, service.ProviderUserService, api.ProviderAuthApi)

	return api.AuthApi{}
}