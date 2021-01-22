package api

import (
	"github.com/labstack/echo/v4"
	"github.com/teten-nugraha/golang-twitter-clone-api/service"
	"net/http"
)

type UserApi struct {
	UserService service.UserService
}

func ProviderUserAPI(u service.UserService) UserApi {
	return UserApi{UserService: u}
}

func (u *UserApi) FindByName(e echo.Context) error {
	username := e.Param("username")

	userDto, err := u.UserService.FindByUsername(username)
	if err != nil {
		return ErrorResponse(e, http.StatusBadRequest,"Terjadi gangguan")
	}

	return SuccessResponse(e, http.StatusOK, userDto)

}