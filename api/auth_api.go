package api

import (
	"github.com/labstack/echo/v4"
	"github.com/teten-nugraha/golang-twitter-clone-api/dto"
	"github.com/teten-nugraha/golang-twitter-clone-api/service"
	"github.com/teten-nugraha/golang-twitter-clone-api/util"
	"net/http"
)

type AuthApi struct {
	UserService service.UserService
}

func ProviderAuthApi(p service.UserService) AuthApi {
	return AuthApi{
		UserService: p,
	}
}

func (p *AuthApi) GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := util.HashPassword(password)

	return SuccessResponse(c, http.StatusOK, hash)
}

func (p *AuthApi) SignUp(c echo.Context) error {
	form := new(dto.SignupDto)
	c.Bind(form)

	signupDto := dto.SignupDto{
		Username: form.Username,
		Email:    form.Email,
		Gender:   form.Gender,
		Password:   form.Password,
	}

	userDto, err := p.UserService.SignUp(signupDto)
	if err != nil {
		return ErrorResponse(c, http.StatusOK, err.Error())
	}

	return SuccessResponse(c, http.StatusOK, userDto)
}
