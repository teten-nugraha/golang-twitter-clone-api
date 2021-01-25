package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/teten-nugraha/golang-twitter-clone-api/domain"
)

func CurrentUserLogin(c echo.Context) domain.User {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	var entity domain.User

	entity.Username = claims["username"].(string)
	entity.Email = claims["email"].(string)

	return entity
}
