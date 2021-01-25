package api

import (
	"github.com/dgrijalva/jwt-go"
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
	if err := c.Bind(form); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, err.Error())
	}


	if err := c.Validate(form); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

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

func(p *AuthApi) CheckLogin(c echo.Context) error {
	form := new(dto.LoginDto)
	c.Bind(form)

	if err := c.Validate(form); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	username := form.Username
	password := form.Password

	user, err := p.UserService.CheckLogin(username, password)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	token, _ := p.UserService.GenerateToken(user)

	return SuccessResponse(c, http.StatusOK, map[string]string{
		"token": token,
	})
}

func (p *AuthApi) WhoAmI(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	return SuccessResponse(c, http.StatusOK, map[string]string{
		"data": "Your login as " + username,
	})
}