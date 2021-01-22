package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/teten-nugraha/golang-twitter-clone-api/api"
)

func UserRoute(routes *echo.Echo, userApi api.UserApi) {
	user := routes.Group("/user")
	user.GET("/findByUsername/:username", userApi.FindByName)
}
